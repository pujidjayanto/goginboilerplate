package product

import (
	"context"
	"testing"

	"github.com/pujidjayanto/goginboilerplate/internal/testutils"
	"github.com/pujidjayanto/goginboilerplate/pkg/jsonb"
	"github.com/pujidjayanto/goginboilerplate/pkg/pagination"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestRepository_FindAll(t *testing.T) {
	db := testutils.NewTestDb(t)

	testutils.WithTransaction(t, db, func(ctx context.Context) {
		repo := NewRepository(db)

		// Create test products
		product1 := createTestProduct(t, ctx, db, "Test Product 1", decimal.NewFromFloat(100.00), 100, jsonb.JSON{})
		product2 := createTestProduct(t, ctx, db, "Test Product 2", decimal.NewFromFloat(250.00), 50, jsonb.JSON{"color": "red"})

		// Test FindAll
		products, err := repo.FindAll(ctx)

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, products, 2)

		// Verify products are returned in order of Id (implicit default order)
		assert.Equal(t, product1.Id, products[0].Id)
		assert.Equal(t, "Test Product 1", products[0].Name)
		assert.True(t, products[0].Price.Equal(decimal.NewFromFloat(100.00)))
		assert.Equal(t, 100, products[0].Quantity)

		assert.Equal(t, product2.Id, products[1].Id)
		assert.Equal(t, "Test Product 2", products[1].Name)
		assert.True(t, products[1].Price.Equal(decimal.NewFromFloat(250.00)))
		assert.Equal(t, 50, products[1].Quantity)
	})
}

func TestRepository_FindById(t *testing.T) {
	db := testutils.NewTestDb(t)

	testutils.WithTransaction(t, db, func(ctx context.Context) {
		repo := NewRepository(db)

		// Create test product
		product := createTestProduct(t, ctx, db, "Test Product", decimal.NewFromFloat(15.75), 75, jsonb.JSON{})

		t.Run("Existing product", func(t *testing.T) {
			foundProduct, err := repo.FindById(ctx, product.Id)

			// Assertions
			assert.NoError(t, err)
			assert.NotNil(t, foundProduct)
			assert.Equal(t, product.Id, foundProduct.Id)
			assert.Equal(t, "Test Product", foundProduct.Name)
			assert.Equal(t, decimal.NewFromFloat(15.75), foundProduct.Price)
			assert.Equal(t, 75, foundProduct.Quantity)
		})

		t.Run("Non-existing product", func(t *testing.T) {
			foundProduct, err := repo.FindById(ctx, 9999)

			// Assertions
			assert.Error(t, err)
			assert.True(t, err == gorm.ErrRecordNotFound)
			assert.Nil(t, foundProduct)
		})
	})
}

func TestRepository_Update(t *testing.T) {
	db := testutils.NewTestDb(t)

	testutils.WithTransaction(t, db, func(ctx context.Context) {
		repo := NewRepository(db)

		// Create test product
		product := createTestProduct(t, ctx, db, "Original Name", decimal.NewFromFloat(10.00), 50, jsonb.JSON{"color": "blue", "material": "cotton"})

		// Update product
		product.Name = "Updated Name"
		product.Price = decimal.NewFromFloat(15.00)
		product.Quantity = 75
		product.ProductDetails = jsonb.JSON{"color": "red", "material": "cotton"}

		err := repo.Update(ctx, product)
		assert.NoError(t, err)

		// Verify update
		updatedProduct, err := repo.FindById(ctx, product.Id)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Name", updatedProduct.Name)
		assert.True(t, updatedProduct.Price.Equal(decimal.NewFromFloat(15.00)))
		assert.Equal(t, 75, updatedProduct.Quantity)
		assert.Equal(t, jsonb.JSON{"color": "red", "material": "cotton"}, updatedProduct.ProductDetails)
	})
}

func TestRepository_FindAllPaginated(t *testing.T) {
	db := testutils.NewTestDb(t)

	testutils.WithTransaction(t, db, func(ctx context.Context) {
		repo := NewRepository(db)

		// Create test products
		createTestProduct(t, ctx, db, "Apple", decimal.NewFromFloat(5.99), 100, jsonb.JSON{})
		createTestProduct(t, ctx, db, "Banana", decimal.NewFromFloat(2.99), 150, jsonb.JSON{})
		createTestProduct(t, ctx, db, "Orange", decimal.NewFromFloat(3.99), 75, jsonb.JSON{})
		createTestProduct(t, ctx, db, "Mango", decimal.NewFromFloat(7.99), 50, jsonb.JSON{})
		createTestProduct(t, ctx, db, "Apple Watch", decimal.NewFromFloat(399.99), 25, jsonb.JSON{})

		t.Run("Basic pagination", func(t *testing.T) {
			filter := ProductFilter{}
			page := pagination.PaginationRequest{
				Page:     1,
				PageSize: 2,
				SortBy:   "name",
				Sort:     "asc",
			}

			products, total, err := repo.FindAllPaginated(ctx, filter, page)
			assert.NoError(t, err)
			assert.Len(t, products, 2)
			assert.Equal(t, int64(5), total)
			assert.Equal(t, "Apple", products[0].Name)
			assert.Equal(t, "Apple Watch", products[1].Name)
		})

		t.Run("Pagination with filter", func(t *testing.T) {
			filter := ProductFilter{
				ProductName: "Apple",
			}
			page := pagination.PaginationRequest{
				Page:     1,
				PageSize: 10,
				SortBy:   "price",
				Sort:     "desc",
			}

			products, total, err := repo.FindAllPaginated(ctx, filter, page)

			assert.NoError(t, err)
			assert.Len(t, products, 2)
			assert.Equal(t, int64(2), total)
			assert.Equal(t, "Apple Watch", products[0].Name)
			assert.Equal(t, "Apple", products[1].Name)
		})

		t.Run("Empty results", func(t *testing.T) {
			filter := ProductFilter{
				ProductName: "NonExistent",
			}
			page := pagination.PaginationRequest{
				Page:     1,
				PageSize: 10,
			}

			products, total, err := repo.FindAllPaginated(ctx, filter, page)

			assert.NoError(t, err)
			assert.Len(t, products, 0)
			assert.Equal(t, int64(0), total)
		})
	})
}

func createTestProduct(t *testing.T, ctx context.Context, db interface {
	GetDB(context.Context) *gorm.DB
}, name string, price decimal.Decimal, quantity int, details jsonb.JSON) *Product {
	product := &Product{
		Name:           name,
		Price:          price,
		Quantity:       quantity,
		ProductDetails: details,
	}

	err := db.GetDB(ctx).Create(product).Error
	require.NoError(t, err)
	require.NotZero(t, product.Id)

	return product
}
