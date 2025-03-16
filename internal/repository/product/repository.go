package product

import (
	"context"

	"github.com/pujidjayanto/goginboilerplate/pkg/db"
	"github.com/pujidjayanto/goginboilerplate/pkg/pagination"
)

type Repository interface {
	FindAll(context.Context) ([]*Product, error)
	FindAllPaginated(context.Context, ProductFilter, pagination.PaginationRequest) ([]*Product, int64, error)
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

func (r *repository) FindAllPaginated(
	ctx context.Context,
	filter ProductFilter,
	page pagination.PaginationRequest,
) ([]*Product, int64, error) {
	var (
		products []*Product
		total    int64
	)

	err := r.db.RunTransaction(ctx, func(innerCtx context.Context) error {
		query := r.db.GetDB(innerCtx)
		if filter.ProductName != "" {
			query = query.Scopes(ProductNameLike(filter.ProductName))
		}

		if err := query.Model(&Product{}).Count(&total).Error; err != nil {
			return err
		}

		if err := query.
			Offset(page.GetOffset()).
			Limit(page.GetLimit()).
			Scopes(Sort(page.GetSortBy(), page.GetSortOrder())).
			Find(&products).
			Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func NewRepository(db db.DatabaseHandler) Repository {
	return &repository{db: db}
}
