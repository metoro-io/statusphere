package dbgroomer

import (
	"context"
	"github.com/metoro-io/statusphere/common/db"
	"github.com/metoro-io/statusphere/common/status_pages"
	"go.uber.org/zap"
)

type DbGroomer struct {
	dbClient *db.DbClient
	logger   *zap.Logger
}

func NewDbGroomer(logger *zap.Logger, dbClient *db.DbClient) *DbGroomer {
	return &DbGroomer{
		logger:   logger,
		dbClient: dbClient,
	}
}

func (d *DbGroomer) Groom() {
	go func() {
		d.logger.Info("grooming status pages")
		// Delete any status page where it isn't in our local list
		statusPages := status_pages.StatusPages
		urls := make(map[string]bool)
		for _, statusPage := range statusPages {
			urls[statusPage.URL] = true
		}

		statusPages, err := d.dbClient.GetAllStatusPages(context.Background())
		if err != nil {
			d.logger.Error("failed to get all status pages", zap.Error(err))
		}
		for _, statusPage := range statusPages {
			if _, ok := urls[statusPage.URL]; !ok {
				d.logger.Info("deleting status page", zap.String("url", statusPage.URL))
				err := d.dbClient.DeleteStatusPage(context.Background(), statusPage.URL)
				if err != nil {
					d.logger.Error("failed to delete status page", zap.Error(err))
				}
			}
		}
		d.logger.Info("finished grooming status pages")
	}()
}
