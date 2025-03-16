package purchase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pujidjayanto/goginboilerplate/internal/dto"
	"github.com/pujidjayanto/goginboilerplate/internal/service/purchase"
	"github.com/pujidjayanto/goginboilerplate/pkg/delivery"
)

type Controller interface {
	Create(*gin.Context)
}

type controller struct {
	purchaseService purchase.Service
}

func (c *controller) Create(ginCtx *gin.Context) {
	var req dto.CreatePurchaseRequest
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		delivery.Failed(ginCtx, http.StatusBadRequest, err.Error())
		return
	}

	userIdAny, exists := ginCtx.Get("userId")
	if !exists {
		delivery.Failed(ginCtx, http.StatusUnauthorized, "user not found")
		return
	}

	req.UserId = userIdAny.(uint)

	err := c.purchaseService.MakePurchase(ginCtx.Request.Context(), req)
	if err != nil {
		delivery.Failed(ginCtx, http.StatusInternalServerError, err.Error())
		return
	}

	delivery.SuccessNoContent(ginCtx)
}

func NewController(purs purchase.Service) Controller {
	return &controller{purchaseService: purs}
}
