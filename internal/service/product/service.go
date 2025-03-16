package product

import (
	"context"
	"fmt"
	"time"

	"github.com/pujidjayanto/goginboilerplate/internal/dto"
	"github.com/pujidjayanto/goginboilerplate/internal/repository/product"
)

type Service interface {
	GetAll(context.Context) (*dto.GetAllProductResponse, error)
	GetAllPaginated(context.Context, dto.GetAllProductRequest) (*dto.GetAllProductPaginatedResponse, error)
}

type service struct {
	productRepository product.Repository
}

func (s *service) GetAll(ctx context.Context) (*dto.GetAllProductResponse, error) {
	productRecords, err := s.productRepository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve products, %v", err)
	}

	products := make([]*dto.ProductItem, 0)
	for _, v := range productRecords {
		products = append(products, &dto.ProductItem{
			Id:       v.ID,
			Name:     v.Name,
			Price:    v.Price.StringFixed(2),
			Quantity: v.Quantity,
		})
	}

	return &dto.GetAllProductResponse{Products: products}, nil
}

func (s *service) GetAllPaginated(ctx context.Context, req dto.GetAllProductRequest) (*dto.GetAllProductPaginatedResponse, error) {
	productRecords, totalRecords, err := s.productRepository.FindAllPaginated(ctx, RequestToFilter(req), req.PaginationRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve products, %v", err)
	}

	products := make([]*dto.ProductItem, 0)
	for _, v := range productRecords {
		products = append(products, &dto.ProductItem{
			Id:        v.ID,
			Name:      v.Name,
			Price:     v.Price.StringFixed(2),
			Quantity:  v.Quantity,
			CreatedAt: v.CreatedAt.Format(time.RFC3339),
		})
	}

	return &dto.GetAllProductPaginatedResponse{
		Products:           products,
		PaginationResponse: req.CreatePaginationResponse(totalRecords, len(products)),
	}, nil
}

func NewService(pr product.Repository) Service {
	return &service{productRepository: pr}
}
