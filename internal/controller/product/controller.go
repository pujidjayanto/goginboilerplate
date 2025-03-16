package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pujidjayanto/goginboilerplate/internal/dto"
	"github.com/pujidjayanto/goginboilerplate/internal/service/product"
	"github.com/pujidjayanto/goginboilerplate/pkg/delivery"
)

type Controller interface {
	// Index get all products without pagination or filter
	Index(*gin.Context)

	// IndexPaginated get all products with filter and pagination
	// This is just demo on how to use filter and pagination
	IndexPaginated(*gin.Context)
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

func (c *controller) IndexPaginated(ginCtx *gin.Context) {
	var req dto.GetAllProductRequest
	if err := ginCtx.ShouldBindQuery(&req); err != nil {
		delivery.Failed(ginCtx, http.StatusBadRequest, err.Error())
		return
	}

	products, err := c.productService.GetAllPaginated(ginCtx.Request.Context(), req)
	if err != nil {
		delivery.Failed(ginCtx, http.StatusInternalServerError, err.Error())
		return
	}

	delivery.SuccessWithMetadata(ginCtx, products.Products, products.PaginationResponse)
}

func NewController(ps product.Service) Controller {
	return &controller{productService: ps}
}
