package dbconsumer

import (
	"context"
	"github.com/metoro-io/statusphere/scraper/api"
	"github.com/metoro-io/statusphere/scraper/internal/db"
	"go.uber.org/zap"
)

type DbConsumer struct {
	logger   *zap.Logger
	dbClient *db.DbClient
}

func NewDbConsumer(logger *zap.Logger, client *db.DbClient) *DbConsumer {
	return &DbConsumer{
		logger:   logger,
		dbClient: client,
	}
}

func (s *DbConsumer) Consume(incidents []api.Incident) error {
	err := s.dbClient.CreateOrUpdateIncidents(context.Background(), incidents)
	if err != nil {
		s.logger.Error("failed to create or update incidents", zap.Error(err))
		return err
	}
	return nil
}
