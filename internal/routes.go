package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/pujidjayanto/goginboilerplate/internal/controller"
)

func setupRouter(g *gin.Engine, controllers controller.Dependency) {
	g.POST("login", controllers.UserController.Login)
}
