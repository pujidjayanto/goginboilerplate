package duck

import (
	"context"

	"github.com/pujidjayanto/goginboilerplate/pkg/db"
)

type Repository interface {
	CreateOne(context.Context) error
}

type repository struct {
	db db.DatabaseHandler
}

func NewRepository(db db.DatabaseHandler) Repository {
	return &repository{db: db}
}

func (r *repository) CreateOne(ctx context.Context) error {
	return nil
}
