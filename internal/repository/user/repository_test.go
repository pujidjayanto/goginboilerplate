package user

import (
	"context"
	"testing"

	"github.com/pujidjayanto/goginboilerplate/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestRepository_Create(t *testing.T) {
	db := testutils.NewTestDb(t)

	testutils.WithTransaction(t, db, func(ctx context.Context) {
		repo := NewRepository(db)

		user := &User{
			Email:        "john@gmail.com",
			PasswordHash: "dummyhashedpassword",
		}

		err := repo.Create(ctx, user)
		assert.NoError(t, err)
		assert.NotZero(t, user.Id)
		assert.Equal(t, "john@gmail.com", user.Email)
	})
}

func TestRepository_FindById(t *testing.T) {
	db := testutils.NewTestDb(t)

	testutils.WithTransaction(t, db, func(ctx context.Context) {
		repo := NewRepository(db)

		// Setup test data
		user := &User{
			Email:        "john@gmail.com",
			PasswordHash: "dummyhashedpassword",
		}

		err := repo.Create(ctx, user)
		require.NoError(t, err)

		// Test cases
		t.Run("get user by id", func(t *testing.T) {
			foundUser, err := repo.FindById(ctx, user.Id)

			assert.NoError(t, err)
			assert.NotNil(t, foundUser)
			assert.Equal(t, user, foundUser)
		})

		t.Run("can't find user by id", func(t *testing.T) {
			foundUser, err := repo.FindById(ctx, 9999)

			assert.Error(t, err)
			assert.True(t, err == gorm.ErrRecordNotFound)
			assert.Nil(t, foundUser)
		})
	})
}

func TestRepository_FindByEmail(t *testing.T) {
	db := testutils.NewTestDb(t)

	testutils.WithTransaction(t, db, func(ctx context.Context) {
		repo := NewRepository(db)

		// Setup test data
		user := &User{
			Email:        "john@gmail.com",
			PasswordHash: "dummyhashedpassword",
		}

		err := repo.Create(ctx, user)
		require.NoError(t, err)

		t.Run("find user by email", func(t *testing.T) {
			foundUser, err := repo.FindByEmail(ctx, user.Email)

			assert.NoError(t, err)
			assert.NotNil(t, foundUser)
			assert.Equal(t, user, foundUser)
		})

		t.Run("can't find user by email", func(t *testing.T) {
			foundUser, err := repo.FindByEmail(ctx, "doe@gmail.com")

			assert.Error(t, err)
			assert.True(t, err == gorm.ErrRecordNotFound)
			assert.Nil(t, foundUser)
		})
	})
}
