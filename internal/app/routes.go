package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/pujidjayanto/goginboilerplate/internal/controller"
	"github.com/pujidjayanto/goginboilerplate/internal/middleware"
)

func setupRouteHandler(controllers controller.Dependency) *gin.Engine {
	gin.DisableConsoleColor()
	g := gin.New()

	// use default cors, but make sure set it properly on real server
	g.Use(cors.Default())
	g.Use(middleware.SecurityHeader())
	g.Use(requestid.New())
	g.Use(middleware.LogRequest())
	g.Use(gin.Recovery())

	g.POST("login", controllers.UserController.Login)
	g.POST("register", controllers.UserController.Register)

	secure := g.Group("/secure")
	secure.Use(middleware.Authenticate())

	productRoutes := secure.Group("/products")
	productRoutes.GET("/", controllers.ProductController.Index)
	productRoutes.GET("/paginated", controllers.ProductController.IndexPaginated)

	purchaseRoutes := secure.Group("/purchases")
	purchaseRoutes.POST("/", controllers.PurchaseController.Create)

	return g
}
