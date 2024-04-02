package incidentpoller

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/metoro-io/statusphere/common/db"
	"github.com/metoro-io/statusphere/common/jobs/slack_webhook"
	"github.com/pkg/errors"
	"github.com/riverqueue/river"
	"go.uber.org/zap"
	"time"
)

// IncidentPoller is a poller that polls incidents from the database, if the jobs have not been started for them, then it starts them

type IncidentPoller struct {
	db              *db.DbClient
	logger          *zap.Logger
	riverClient     *river.Client[pgx.Tx]
	slackWebhookUrl string
}

func NewIncidentPoller(db *db.DbClient, logger *zap.Logger, client *river.Client[pgx.Tx], webhookUrl string) *IncidentPoller {
	return &IncidentPoller{
		db:              db,
		logger:          logger,
		riverClient:     client,
		slackWebhookUrl: webhookUrl,
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
			p.logger.Error("failed to poll", zap.Error(err))
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
	for _, incident := range incidents {
		if p.slackWebhookUrl == "" {
			continue
		}
		jobArgs = append(jobArgs, river.InsertManyParams{Args: slack_webhook.SlackWebhookArgs{
			Incident:   incident,
			WebhookUrl: p.slackWebhookUrl,
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
