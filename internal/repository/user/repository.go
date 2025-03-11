package user

import (
	"context"

	"github.com/pujidjayanto/goginboilerplate/pkg/db"
)

type Repository interface {
	Create(context.Context, *User) error
	FindByID(context.Context, uint) (*User, error)
}

type repository struct {
	db db.DatabaseHandler
}

func (r *repository) Create(ctx context.Context, user *User) error {
	return r.db.GetDB(ctx).Create(user).Error
}

func (r *repository) FindByID(ctx context.Context, id uint) (*User, error) {
	var user User
	if err := r.db.GetDB(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewRepository(db db.DatabaseHandler) Repository {
	return &repository{db: db}
}
