package incidentpoller

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/metoro-io/statusphere/common/api"
	"github.com/metoro-io/statusphere/common/db"
	"github.com/metoro-io/statusphere/common/jobs/slack_webhook"
	"github.com/metoro-io/statusphere/common/jobs/twitter_post"
	"github.com/pkg/errors"
	"github.com/riverqueue/river"
	"go.uber.org/zap"
	"time"
)

// IncidentPoller is a poller that polls incidents from the database, if the jobs have not been started for them, then it starts them

type IncidentPoller struct {
	db                *db.DbClient
	logger            *zap.Logger
	riverClient       *river.Client[pgx.Tx]
	slackWebhookUrl   string
	twitterWebhookUrl string
}

func NewIncidentPoller(db *db.DbClient, logger *zap.Logger, client *river.Client[pgx.Tx], slackWebhookUrl string, twitterWebhookUrl string) *IncidentPoller {
	return &IncidentPoller{
		db:                db,
		logger:            logger,
		riverClient:       client,
		slackWebhookUrl:   slackWebhookUrl,
		twitterWebhookUrl: twitterWebhookUrl,
	}
}

func (p *IncidentPoller) Start() {
	go p.Poll()
}

func (p *IncidentPoller) Poll() error {
	err := p.pollInner()
	if err != nil {
		p.logger.Error("failed to poll", zap.Error(err))
	}
	ticker := time.NewTicker(1 * time.Minute)
	for {
		select {
		case <-ticker.C:
			err := p.pollInner()
			if err != nil {
				p.logger.Error("failed to poll", zap.Error(err))
			}
		}
	}
}

func (p *IncidentPoller) pollInner() error {
	// Get the incidents from the database without jobs started
	p.logger.Info("polling incidents without jobs started")
	incidents, err := p.db.GetIncidentsWithoutJobsStarted(context.Background(), 1000)
	if err != nil {
		return errors.Wrap(err, "failed to get incidents without jobs started")
	}
	if len(incidents) == 0 {
		p.logger.Info("no incidents without jobs started")
		return nil
	}
	p.logger.Info("found incidents without jobs started", zap.Int("count", len(incidents)))

	jobArgs := make([]river.InsertManyParams, 0)

	var incidentsToProcess = make([]api.Incident, 0)
	for _, incident := range incidents {
		if p.slackWebhookUrl == "" {
			continue
		}

		// We only want to notify about incidents that have started in the last hour
		// Otherwise, we will be sending notifications for incidents that have already been resolved
		if incident.StartTime.Before(time.Now().Add(-1 * time.Hour)) {
			continue
		}

		if incident.Impact == "maintenance" {
			continue
		}

		incidentsToProcess = append(incidentsToProcess, incident)
	}

	// Slack webhook notifications
	for _, incident := range incidentsToProcess {
		jobArgs = append(jobArgs, river.InsertManyParams{Args: slack_webhook.SlackWebhookArgs{
			Incident:   incident,
			WebhookUrl: p.slackWebhookUrl,
		}})
	}

	// Twitter post notifications
	for _, incident := range incidentsToProcess {
		jobArgs = append(jobArgs, river.InsertManyParams{Args: twitter_post.TwitterPostArgs{
			WebhookUrl: p.twitterWebhookUrl,
			Incident:   incident,
		}})
	}

	p.logger.Info("starting to insert jobs", zap.Int("count", len(jobArgs)))
	// Start the jobs for each incident and update the database
	if len(jobArgs) != 0 {
		_, err = p.riverClient.InsertMany(context.Background(), jobArgs)
		if err != nil {
			return errors.Wrap(err, "failed to insert many")
		}
	}
	p.logger.Info("finished inserting jobs")

	// Start the jobs for each incident and update the database
	return p.db.SetIncidentNotificationStartedToTrue(context.Background(), incidents)
}
