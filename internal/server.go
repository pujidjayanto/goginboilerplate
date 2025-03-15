package internal

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pujidjayanto/goginboilerplate/internal/controller"
	"github.com/pujidjayanto/goginboilerplate/internal/repository"
	"github.com/pujidjayanto/goginboilerplate/internal/service"
	"github.com/pujidjayanto/goginboilerplate/pkg/db"
)

func NewApplicationServer(db db.DatabaseHandler) *http.Server {
	repositories := repository.NewDependency(db)
	services := service.NewDependency(repositories, db)
	controllers := controller.NewDependency(services)

	switch GlobalConfig.Server.Env {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	ginEngine := gin.Default()
	setupRouter(ginEngine, controllers)

	return &http.Server{
		Addr:         GlobalConfig.ServerPort(),
		Handler:      ginEngine,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
