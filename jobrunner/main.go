package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/metoro-io/statusphere/common/api"
	"github.com/metoro-io/statusphere/common/db"
	"github.com/metoro-io/statusphere/common/jobs/riverclient"
	"github.com/metoro-io/statusphere/common/jobs/slack_webhook"
	"github.com/riverqueue/river"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	db, err := db.NewDbClientFromEnvironment(logger)
	if err != nil {
		panic(err)
	}

	err = riverclient.RunMigration(db.PgxPool)

	client, err := riverclient.NewRiverClient(db.PgxPool, logger, http.DefaultClient, 100)
	if err != nil {
		panic(err)
	}
	err = client.Start(ctx)
	if err != nil {
		panic(err)
	}
	defer func(client *river.Client[pgx.Tx], ctx context.Context) {
		_ = client.Stop(ctx)
	}(client, ctx)

	_, err = client.Insert(ctx, &slack_webhook.SlackWebhookArgs{
		WebhookUrl: "",
		Incident:   api.Incident{},
	}, nil)
	if err != nil {
		panic(err)
	}

	// Work forever
	<-ctx.Done()
}
