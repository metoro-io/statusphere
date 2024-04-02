package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/metoro-io/statusphere/common/db"
	"github.com/metoro-io/statusphere/common/jobs/riverclient"
	"github.com/metoro-io/statusphere/jobrunner/internal/incidentpoller"
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

	incidentPoller := incidentpoller.NewIncidentPoller(db, logger, client, "")
	incidentPoller.Start()

	// Work forever
	<-ctx.Done()
}
