package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pujidjayanto/goginboilerplate/internal/delivery"
	"github.com/pujidjayanto/goginboilerplate/internal/service/product"
)

type Controller interface {
	Index(*gin.Context)
}

type controller struct {
	productService product.Service
}

func (c *controller) Index(ginCtx *gin.Context) {
	products, err := c.productService.GetAll(ginCtx.Request.Context())
	if err != nil {
		delivery.Failed(ginCtx, http.StatusInternalServerError, err.Error())
		return
	}

	delivery.Success(ginCtx, products)
}

func NewController(ps product.Service) Controller {
	return &controller{productService: ps}
}
