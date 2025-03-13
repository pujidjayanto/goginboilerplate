package user

import (
	"context"
	"testing"

	"github.com/pujidjayanto/goginboilerplate/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Create(t *testing.T) {
	tests := []struct {
		name      string
		params    *User
		setup     func(Repository)
		wantError bool
	}{
		{
			name: "valid create",
			params: &User{
				Email:        "john@gmail.com",
				PasswordHash: "hashedPassword",
			},
			wantError: false,
		},
		{
			name: "same email, return error",
			params: &User{
				Email:        "john@gmail.com",
				PasswordHash: "hashedPassword",
			},
			setup: func(r Repository) {
				r.Create(context.Background(), &User{
					Email:        "john@gmail.com",
					PasswordHash: "hashedPassword",
				})
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := testutils.NewTestDb(t)
			userRepository := NewRepository(db)

			if tt.setup != nil {
				tt.setup(userRepository)
			}

			err := userRepository.Create(context.Background(), tt.params)
			if !tt.wantError {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestRepository_FindByID(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(Repository)
		id        uint
		wantError bool
	}{
		{
			name: "find existing user",
			setup: func(r Repository) {
				r.Create(context.Background(), &User{
					Email:        "existing@gmail.com",
					PasswordHash: "hashedPassword",
				})
			},
			id:        1,
			wantError: false,
		},
		{
			name:      "find non-existent user",
			id:        999,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := testutils.NewTestDb(t)
			userRepository := NewRepository(db)

			if tt.setup != nil {
				tt.setup(userRepository)
			}

			user, err := userRepository.FindById(context.Background(), tt.id)
			if tt.wantError {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, "existing@gmail.com", user.Email)
			}
		})
	}
}

func TestRepository_FindByEmail(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(Repository)
		email     string
		wantError bool
	}{
		{
			name: "find existing user by email",
			setup: func(r Repository) {
				err := r.Create(context.Background(), &User{
					Email:        "existing@gmail.com",
					PasswordHash: "hashedPassword",
				})
				assert.NoError(t, err)
			},
			email:     "existing@gmail.com",
			wantError: false,
		},
		{
			name:      "find non-existent user by email",
			email:     "nonexistent@gmail.com",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := testutils.NewTestDb(t)
			userRepository := NewRepository(db)

			if tt.setup != nil {
				tt.setup(userRepository)
			}

			user, err := userRepository.FindByEmail(context.Background(), tt.email)
			if tt.wantError {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, "existing@gmail.com", user.Email)
			}
		})
	}
}
