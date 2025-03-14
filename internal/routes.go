package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/pujidjayanto/goginboilerplate/internal/controller"
	"github.com/pujidjayanto/goginboilerplate/internal/middleware"
)

func setupRouter(g *gin.Engine, controllers controller.Dependency) {
	g.POST("login", controllers.UserController.Login)
	g.POST("register", controllers.UserController.Register)

	secure := g.Group("/secure")
	secure.Use(middleware.Authenticate())

	productRoutes := secure.Group("/products")
	productRoutes.GET("/", controllers.ProductController.Index)
}
