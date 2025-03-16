package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pujidjayanto/goginboilerplate/internal/dto"
	"github.com/pujidjayanto/goginboilerplate/internal/service/user"
	"github.com/pujidjayanto/goginboilerplate/pkg/delivery"
)

type Controller interface {
	Login(*gin.Context)
	Register(*gin.Context)
}

type controller struct {
	userService user.Service
}

func (c *controller) Login(ginCtx *gin.Context) {
	var req dto.LoginRequest
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		delivery.Failed(ginCtx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := c.userService.Login(ginCtx.Request.Context(), req)
	if err != nil {
		if errors.Is(err, user.ErrInvalidCredential) || errors.Is(err, user.ErrUserNotFound) {
			delivery.Failed(ginCtx, http.StatusUnauthorized, err.Error())
			return
		}

		delivery.Failed(ginCtx, http.StatusInternalServerError, err.Error())
		return
	}

	delivery.Success(ginCtx, token)
}

func (c *controller) Register(ginCtx *gin.Context) {
	var req dto.RegisterRequest
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		delivery.Failed(ginCtx, http.StatusBadRequest, err.Error())
		return
	}

	err := c.userService.Register(ginCtx.Request.Context(), req)
	if err != nil {
		if errors.Is(err, user.ErrUserAlreadyExisted) {
			delivery.Failed(ginCtx, http.StatusConflict, err.Error())
			return
		}

		delivery.Failed(ginCtx, http.StatusInternalServerError, err.Error())
		return
	}

	delivery.SuccessCreated(ginCtx)
}

func NewController(uss user.Service) Controller {
	return &controller{userService: uss}
}
