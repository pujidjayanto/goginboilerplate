package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pujidjayanto/goginboilerplate/internal/config"
	"github.com/pujidjayanto/goginboilerplate/internal/controller"
	"github.com/pujidjayanto/goginboilerplate/internal/repository"
	"github.com/pujidjayanto/goginboilerplate/internal/service"
	"github.com/pujidjayanto/goginboilerplate/pkg/db"
)

func NewApplicationServer(db db.DatabaseHandler) *http.Server {
	repositories := repository.NewDependency(db)
	services := service.NewDependency(repositories, db)
	controllers := controller.NewDependency(services)

	switch config.GetEnv() {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	handler := setupRouteHandler(controllers)

	return &http.Server{
		Addr:         config.GetPort(),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
