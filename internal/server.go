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

type ServerDependency struct {
	Env, Port string
	Database  db.DatabaseHandler
}

func NewApplicationServer(dependency *ServerDependency) *http.Server {
	repositories := repository.NewDependency(dependency.Database)
	services := service.NewDependency(repositories)
	controllers := controller.NewDependency(services)

	switch dependency.Env {
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
		Addr:         dependency.Port,
		Handler:      ginEngine,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
