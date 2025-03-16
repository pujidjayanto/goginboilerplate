package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/pujidjayanto/goginboilerplate/internal/app"
	"github.com/pujidjayanto/goginboilerplate/internal/config"
	"github.com/pujidjayanto/goginboilerplate/pkg/db"
	"github.com/pujidjayanto/goginboilerplate/pkg/logger"
)

func main() {
	logger.Init()
	defer logger.SyncLogger()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	err := config.Initialize()
	if err != nil {
		logger.Fatal("initialize configuration error", "err", err.Error())
	}

	db, err := db.InitDatabaseHandler(config.GetDatabaseDSN())
	if err != nil {
		logger.Fatal("initialize database error", "err", err.Error())
	}

	httpServer := app.NewApplicationServer(db)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("server listened failed", "err", err.Error())
		}
	}()

	<-ctx.Done()

	stop()
	logger.Info("shutting down gracefully, press Ctrl+C again to force")

	if err := db.Close(); err != nil {
		logger.Fatal("error closing connection", "err", err.Error())
	}

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Fatal("server forced to shutdown", "err", err.Error())
	}

	logger.Info("server exiting")
}
