package product

import (
	"github.com/pujidjayanto/goginboilerplate/internal/dto"
	"github.com/pujidjayanto/goginboilerplate/internal/repository/product"
)

func RequestToFilter(req dto.GetAllProductRequest) product.ProductFilter {
	return product.ProductFilter{ProductName: req.ProductName}
}
