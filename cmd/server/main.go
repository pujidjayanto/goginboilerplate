package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/pujidjayanto/goginboilerplate/internal/app"
	"github.com/pujidjayanto/goginboilerplate/internal/config"
	"github.com/pujidjayanto/goginboilerplate/pkg/db"
	"github.com/pujidjayanto/goginboilerplate/pkg/log"
	"go.uber.org/zap"
)

func main() {
	log.Init()
	defer log.SyncLogger()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	err := config.LoadConfiguration()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.ConfigureLogger(config.GlobalConfig.Server.Env)

	db, err := db.InitDatabaseHandler(config.GlobalConfig.DatabaseDSN())
	if err != nil {
		log.Fatal(err.Error())
	}

	httpServer := app.NewApplicationServer(db)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: %s\n", zap.String("error", err.Error()))
		}
	}()

	<-ctx.Done()

	stop()
	log.Info("shutting down gracefully, press Ctrl+C again to force")

	if err := db.Close(); err != nil {
		log.Fatal("error closing connection", zap.String("error", err.Error()))
	}

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown: ", zap.String("error", err.Error()))
	}

	log.Info("server exiting")
}
