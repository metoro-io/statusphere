package slack_webhook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/riverqueue/river"
	"go.uber.org/zap"
	"math"
	"net/http"
	"time"
)

type SlackWebhookWorker struct {
	// An embedded WorkerDefaults sets up default methods to fulfill the rest of
	// the Worker interface:
	river.WorkerDefaults[SlackWebhookArgs]
	logger     *zap.Logger
	httpClient *http.Client
}

func NewSlackWebhookWorker(logger *zap.Logger, httpClient *http.Client) *SlackWebhookWorker {
	return &SlackWebhookWorker{
		logger:     logger,
		httpClient: httpClient,
	}
}

func (w *SlackWebhookWorker) Work(ctx context.Context, job *river.Job[SlackWebhookArgs]) error {
	w.logger.Info("Sending slack webhook", zap.Any("incident", job.Args.Incident))
	marshal, err := json.Marshal(job.Args.Incident)
	if err != nil {
		return errors.Wrap(err, "failed to marshal incident")
	}
	postBody := fmt.Sprintf(`{'text': '%s'}`, string(marshal))
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
		return errors.Errorf("expected status code 200, got %d", resp.StatusCode)
	}
	return nil
}

func (w *SlackWebhookWorker) Timeout(job *river.Job[SlackWebhookArgs]) time.Duration {
	return time.Minute * 5
}

// NextRetry performs exponential backoff with a maximum of 120 seconds.
func (w *SlackWebhookWorker) NextRetry(job *river.Job[SlackWebhookArgs]) time.Time {
	return time.Now().Add(time.Duration(math.Min(math.Pow(2.0, float64(job.Attempt)), 120)) * time.Second)
}
