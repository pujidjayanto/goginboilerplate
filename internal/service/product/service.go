package product

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pujidjayanto/goginboilerplate/internal/dto"
	"github.com/pujidjayanto/goginboilerplate/internal/repository/product"
)

type Service interface {
	GetAll(context.Context) (*dto.GetAllProduct, error)
}

type service struct {
	productRepository product.Repository
}

func (s *service) GetAll(ctx context.Context) (*dto.GetAllProduct, error) {
	productRecords, err := s.productRepository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve products, %v", err)
	}

	products := make([]*dto.ProductItem, 0)
	for _, v := range productRecords {
		products = append(products, &dto.ProductItem{
			Id:       v.ID,
			Name:     v.Name,
			Price:    strconv.FormatFloat(v.Price, 'f', -1, 64),
			Quantity: v.Quantity,
		})
	}

	return &dto.GetAllProduct{Products: products}, nil
}

func NewService(pr product.Repository) Service {
	return &service{productRepository: pr}
}
