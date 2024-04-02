package riverclient

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/metoro-io/statusphere/common/jobs/slack_webhook"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
	"github.com/riverqueue/river/rivermigrate"
	"go.uber.org/zap"
	"net/http"
)

func spawnWorkers(logger *zap.Logger, client *http.Client) *river.Workers {
	workers := river.NewWorkers()
	river.AddWorker(workers, slack_webhook.NewSlackWebhookWorker(logger, client))
	return workers
}

func NewRiverClient(pool *pgxpool.Pool, logger *zap.Logger, client *http.Client, numWorkers int) (*river.Client[pgx.Tx], error) {
	return river.NewClient(riverpgxv5.New(pool), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {
				MaxWorkers: numWorkers,
			},
		},
		Workers: spawnWorkers(logger, client),
	})
}

func RunMigration(pool *pgxpool.Pool) error {
	migrator := rivermigrate.New(riverpgxv5.New(pool), nil)
	_, err := migrator.Migrate(context.Background(), rivermigrate.DirectionUp, nil)
	return err
}
