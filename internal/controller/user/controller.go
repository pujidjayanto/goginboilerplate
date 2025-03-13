package user

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pujidjayanto/goginboilerplate/internal/dto"
	"github.com/pujidjayanto/goginboilerplate/internal/service/user"
)

type Controller interface {
	Login(*gin.Context)
}

type controller struct {
	userService user.Service
}

func (c *controller) Login(ginCtx *gin.Context) {
	var req dto.LoginRequest
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := c.userService.Login(ginCtx.Request.Context(), req)
	if err != nil {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, dto.LoginResponse{Token: token})
}

func NewController(uss user.Service) Controller {
	return &controller{userService: uss}
}

func (c *controller) Create(ctx context.Context) error {
	return nil
}
