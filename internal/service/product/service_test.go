package product

import (
	"context"
	"testing"

	"github.com/pujidjayanto/goginboilerplate/internal/repository/product"
	mockProductRepository "github.com/pujidjayanto/goginboilerplate/mocks/internal_/repository/product"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestService_GetAll(t *testing.T) {
	productRepository := mockProductRepository.NewRepository(t)
	productService := NewService(productRepository)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockProducts := []*product.Product{
			{
				Id:       1,
				Name:     "Honda",
				Price:    decimal.NewFromFloat(10.99),
				Quantity: 5,
			},
			{
				Id:       2,
				Name:     "Hyundai",
				Price:    decimal.NewFromFloat(20.50),
				Quantity: 10,
			},
		}

		productRepository.EXPECT().FindAll(ctx).Return(mockProducts, nil)

		result, err := productService.GetAll(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Products, 2)
		assert.Equal(t, "10.99", result.Products[0].Price)
		assert.Equal(t, "20.50", result.Products[1].Price)
	})
}
