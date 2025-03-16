package dto

import "github.com/pujidjayanto/goginboilerplate/pkg/pagination"

type GetAllProductRequest struct {
	ProductName string `form:"productName" binding:"omitempty"`

	pagination.PaginationRequest
}

type ProductItem struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"createdAt"`
}

type GetAllProductResponse struct {
	Products []*ProductItem `json:"products"`
}

type GetAllProductPaginatedResponse struct {
	Products []*ProductItem `json:"products"`

	pagination.PaginationResponse
}
