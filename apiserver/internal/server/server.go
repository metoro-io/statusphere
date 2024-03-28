package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/metoro-io/statusphere/common/db"
	"github.com/metoro-io/statusphere/common/utils"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

type Server struct {
	logger               *zap.Logger
	dbClient             *db.DbClient
	statusPageCache      *cache.Cache
	incidentCache        *cache.Cache
	currentIncidentCache *cache.Cache
}

func NewServer(logger *zap.Logger, dbClient *db.DbClient) *Server {
	return &Server{
		logger:               logger,
		dbClient:             dbClient,
		statusPageCache:      cache.New(15*time.Minute, 15*time.Minute),
		incidentCache:        cache.New(1*time.Minute, 1*time.Minute),
		currentIncidentCache: cache.New(1*time.Minute, 1*time.Minute),
	}
}

func (s *Server) Serve() error {
	r := gin.New()
	r.UseH2C = true
	r.Use(gin.Recovery())

	corsHandler := handleCors()
	r.Use(corsHandler)
	r.Use(gzip.Gzip(gzip.BestSpeed))

	r.Use(ginZap(s.logger))

	apiV1 := r.Group("/api/v1")
	{
		apiV1.Use(addNoIndexHeader())
		apiV1.GET("/incidents", s.incidents)
		apiV1.GET("/currentStatus", s.currentStatus)
		apiV1.GET("/statusPage", s.statusPage)
		apiV1.GET("/statusPages", s.statusPages)
		apiV1.GET("/statusPages/search", s.statusPageSearch)
		apiV1.GET("/statusPages/count", s.statusPageCount)
	}

	s.addFrontendRoutes(r)

	return errors.Wrap(r.Run(":80"), "Failed to start server")
}

func (s *Server) addFrontendRoutes(r *gin.Engine) {
	r.Static("/static", "/etc/frontend/static")
	r.StaticFile("/asset-manifest.json", "/etc/frontend/asset-manifest.json")
	r.StaticFile("/favicon.ico", "/etc/frontend/favicon.ico")
	r.StaticFile("robots.txt", "/etc/frontend/robots.txt")
	r.NoRoute(func(c *gin.Context) {
		c.File("/etc/frontend/index.html")
	})
}

func handleCors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	// Development cors
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "https://metoro.io"}
	handlerFunc := cors.New(corsConfig)
	return handlerFunc
}

// Middleware to make Gin log using Zap
func ginZap(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Collect log fields
		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		if len(c.Errors) > 0 {
			// Add any error messages to the log
			utils.GetLogger(ctx, logger).Error(c.Errors.String(),
				zap.Int("status", statusCode),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("query", query),
				zap.Duration("latency", latency),
				zap.String("clientIP", clientIP),
				zap.String("user-agent", c.Request.UserAgent()),
			)
		} else {
			// Log normally
			utils.GetLogger(ctx, logger).Info(path,
				zap.Int("status", statusCode),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("query", query),
				zap.Duration("latency", latency),
				zap.String("clientIP", clientIP),
				zap.String("user-agent", c.Request.UserAgent()),
			)
		}
	}
}

func addNoIndexHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Robots-Tag", "noindex")
		c.Next()
	}
}
