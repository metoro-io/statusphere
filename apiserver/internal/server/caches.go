package server

import (
	"context"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"time"
)

func (s *Server) StartCaches(ctx context.Context) {
	go s.updateStatusPageCache(ctx)
}

const statusPageCacheRefreshInterval = 5 * time.Minute

func (s *Server) updateStatusPageCache(ctx context.Context) {
	ticker := time.NewTicker(statusPageCacheRefreshInterval)
	s.updateStatusPageCacheInner(ctx)
	for {
		select {
		case <-ticker.C:
			s.updateStatusPageCacheInner(ctx)
		}
	}
}

func (s *Server) updateStatusPageCacheInner(ctx context.Context) {
	statusPages, err := s.dbClient.GetAllStatusPages(ctx)
	if err != nil {
		s.logger.Error("failed to get status pages", zap.Error(err))
		return
	}

	for _, statusPage := range statusPages {
		s.statusPageCache.Set(statusPage.URL, statusPage, cache.DefaultExpiration)
	}
}
