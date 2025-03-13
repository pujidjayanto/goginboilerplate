package user

import (
	"context"

	"github.com/pujidjayanto/goginboilerplate/pkg/db"
)

type Repository interface {
	Create(context.Context, *User) error
	FindById(context.Context, uint) (*User, error)
	FindByEmail(context.Context, string) (*User, error)
}

type repository struct {
	db db.DatabaseHandler
}

func (r *repository) Create(ctx context.Context, user *User) error {
	return r.db.GetDB(ctx).Create(user).Error
}

func (r *repository) FindById(ctx context.Context, id uint) (*User, error) {
	var user User
	if err := r.db.GetDB(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	if err := r.db.GetDB(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewRepository(db db.DatabaseHandler) Repository {
	return &repository{db: db}
}
