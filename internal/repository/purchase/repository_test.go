package purchase

import (
	"context"
	"testing"

	"github.com/pujidjayanto/goginboilerplate/internal/repository/product"
	"github.com/pujidjayanto/goginboilerplate/internal/repository/user"
	"github.com/pujidjayanto/goginboilerplate/internal/testutils"
	"github.com/pujidjayanto/goginboilerplate/pkg/jsonb"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// todo: refactor later
func TestRepository_Create(t *testing.T) {
	db := testutils.NewTestDb(t)

	testutils.WithTransaction(t, db, func(ctx context.Context) {
		purchaseRepository := NewRepository(db)
		userRepository := user.NewRepository(db)

		// Setup test data
		testUser := &user.User{
			Email:        "john@gmail.com",
			PasswordHash: "dummyhashedpassword",
		}

		err := userRepository.Create(ctx, testUser)
		require.NoError(t, err)

		testProduct := createTestProduct(t, ctx, db, "Test Product 1", decimal.NewFromFloat(100.00), 100, jsonb.JSON{})

		purchase := Purchase{
			UserId:    testUser.Id,
			ProductId: testProduct.Id,
			Quantity:  2,
		}

		err = purchaseRepository.Create(ctx, purchase)
		assert.NoError(t, err)
	})
}

func createTestProduct(t *testing.T, ctx context.Context, db interface {
	GetDB(context.Context) *gorm.DB
}, name string, price decimal.Decimal, quantity int, details jsonb.JSON) *product.Product {
	product := &product.Product{
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
