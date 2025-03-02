package duck

import (
	"context"

	"github.com/pujidjayanto/goginboilerplate/internal/repository/duck"
)

type Service interface {
	CreateOne(context.Context) error
}

type service struct {
	duckRepository duck.Repository
}

func NewService(duckRepository duck.Repository) Service {
	return &service{duckRepository: duckRepository}
}

func (r *service) CreateOne(ctx context.Context) error {
	return nil
}
