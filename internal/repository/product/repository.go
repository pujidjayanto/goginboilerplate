package product

import (
	"context"

	"github.com/pujidjayanto/goginboilerplate/pkg/db"
)

type Repository interface {
	FindAll(context.Context) ([]*Product, error)
	FindById(context.Context, uint) (*Product, error)
	Update(context.Context, *Product) error
}

type repository struct {
	db db.DatabaseHandler
}

func (r *repository) FindAll(ctx context.Context) ([]*Product, error) {
	var products []*Product
	if err := r.db.GetDB(ctx).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) FindById(ctx context.Context, id uint) (*Product, error) {
	var product Product
	if err := r.db.GetDB(ctx).First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repository) Update(ctx context.Context, product *Product) error {
	if err := r.db.GetDB(ctx).Save(product).Error; err != nil {
		return err
	}

	return nil
}

func NewRepository(db db.DatabaseHandler) Repository {
	return &repository{db: db}
}
