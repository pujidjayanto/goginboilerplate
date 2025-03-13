package purchase

import (
	"context"

	"github.com/pujidjayanto/goginboilerplate/pkg/db"
)

type Repository interface {
	Create(context.Context, Purchase) error
}

type repository struct {
	db db.DatabaseHandler
}

func (r *repository) Create(ctx context.Context, purchase Purchase) error {
	return r.db.GetDB(ctx).Create(&purchase).Error
}

func NewRepository(db db.DatabaseHandler) Repository {
	return &repository{db: db}
}
