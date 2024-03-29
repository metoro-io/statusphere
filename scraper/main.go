package main

import (
	"context"
	"github.com/metoro-io/statusphere/common/db"
	"github.com/metoro-io/statusphere/scraper/internal/scraper"
	"github.com/metoro-io/statusphere/scraper/internal/scraper/consumers"
	"github.com/metoro-io/statusphere/scraper/internal/scraper/consumers/dbconsumer"
	"github.com/metoro-io/statusphere/scraper/internal/scraper/poller"
	"github.com/metoro-io/statusphere/scraper/internal/scraper/providers"
	"github.com/metoro-io/statusphere/scraper/internal/scraper/urlgetter/dburlgetter"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	scraper := scraper.NewScraper(logger, http.DefaultClient, []providers.Provider{
		atlassian.NewAtlassianProvider(logger, http.DefaultClient),
	})

	dbClient, err := db.NewDbClientFromEnvironment(logger)
	if err != nil {
		logger.Error("failed to create db client", zap.Error(err))
		return
	}

	err = dbClient.AutoMigrate(context.Background())
	if err != nil {
		logger.Error("failed to auto migrate", zap.Error(err))
		return
	}

	getter := dburlgetter.NewDBURLGetter(logger, dbClient)
	getter.Start()
	poller := poller.NewPoller(getter, scraper, []consumers.Consumer{
		dbconsumer.NewDbConsumer(logger, dbClient),
	}, logger)
	err = poller.Poll()
	if err != nil {
		logger.Error("failed to poll", zap.Error(err))
		return
	}
}
