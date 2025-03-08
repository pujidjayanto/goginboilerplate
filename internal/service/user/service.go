package user

import (
	"context"

	"github.com/pujidjayanto/goginboilerplate/internal/repository/user"
)

type Service interface {
	CreateOne(context.Context) error
}

type service struct {
	userRepository user.Repository
}

func NewService(ur user.Repository) Service {
	return &service{userRepository: ur}
}

func (r *service) CreateOne(ctx context.Context) error {
	return nil
}
