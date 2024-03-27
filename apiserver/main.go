package main

import (
	"context"
	"errors"
	"github.com/metoro-io/statusphere/apiserver/internal/server"
	"github.com/metoro-io/statusphere/common/db"
	"github.com/metoro-io/statusphere/common/utils"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	dbClient, err := db.NewDbClientFromEnvironment(logger)
	if err != nil {
		panic(err)
	}

	s := server.NewServer(logger, dbClient)
	s.StartCaches(ctx)

	go func() {
		if err := s.Serve(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			utils.GetLogger(ctx, logger).Fatal("Failed to start server", zap.Error(err))
		}
		logger.Info("Server started")
	}()

	// Listen for shutdown signal, then cancel the context
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	logger.Info("Shutting down", zap.String("signal", sig.String()))

}
