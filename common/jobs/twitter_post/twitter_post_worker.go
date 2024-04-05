package twitter_post

import (
	"bytes"
	"context"
	"fmt"
	"github.com/metoro-io/statusphere/common/api"
	"github.com/metoro-io/statusphere/common/db"
	"github.com/pkg/errors"
	"github.com/riverqueue/river"
	"go.uber.org/zap"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

type TwitterPostWorker struct {
	// An embedded WorkerDefaults sets up default methods to fulfill the rest of
	// the Worker interface:
	river.WorkerDefaults[TwitterPostArgs]
	logger     *zap.Logger
	httpClient *http.Client
	db         *db.DbClient
}

func NewTwitterPostWorker(logger *zap.Logger, httpClient *http.Client, dbClient *db.DbClient) *TwitterPostWorker {
	return &TwitterPostWorker{
		logger:     logger,
		httpClient: httpClient,
		db:         dbClient,
	}
}

func (w *TwitterPostWorker) Work(ctx context.Context, job *river.Job[TwitterPostArgs]) error {
	w.logger.Info("Sending slack webhook", zap.Any("incident", job.Args.Incident))
	if job.Args.WebhookUrl == "" {
		w.logger.Error("webhook url is empty")
		return nil
	}
	tweet, err := generateTweet(w.db, job.Args.Incident)
	if err != nil {
		return errors.Wrap(err, "failed to generate tweet")
	}
	postBody := fmt.Sprintf(`{"tweet": "%s"}`, string(tweet))
	req, err := http.NewRequest("POST", job.Args.WebhookUrl, bytes.NewBuffer([]byte(postBody)))
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := w.httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}
	if resp.StatusCode != http.StatusOK {
		all, _ := ioutil.ReadAll(resp.Body)
		w.logger.Error("body", zap.Any("body", all))
		return errors.Errorf("expected status code 200, got %d", resp.StatusCode)
	}
	return nil
}

func generateTweet(db *db.DbClient, incident api.Incident) (string, error) {
	// Get the status page of the incident
	statusPage, err := db.GetStatusPage(context.Background(), incident.StatusPageUrl)
	if err != nil {
		return "", errors.Wrap(err, "failed to get status page")
	}
	// Tweet format
	// {Status page Name} Incident
	// {Incident Title}
	// {Incident Description}
	// {Incident Deep Link}
	// https://metoro.io/statusphere/status/{statusPageName}

	if incident.Description == nil {
		incident.Description = new(string)
	}

	tweet := fmt.Sprintf(`🔥 %s Incident - Is %s down? 🔥\r\rTitle: %s\r\rDescription: %s\r\rIncident Deeplink: %s\r\rStatusphere: https://metoro.io/statusphere/status/%s\r\r#outage #incident`, statusPage.Name, statusPage.Name, incident.Title, *incident.Description, incident.DeepLink, statusPage.Name)

	return tweet, nil
}

func (w *TwitterPostWorker) Timeout(job *river.Job[TwitterPostArgs]) time.Duration {
	return time.Second * 30
}

// NextRetry performs exponential backoff with a maximum of 120 seconds.
func (w *TwitterPostWorker) NextRetry(job *river.Job[TwitterPostArgs]) time.Time {
	return time.Now().Add(time.Duration(math.Min(math.Pow(2.0, float64(job.Attempt)), 120)) * time.Second)
}
