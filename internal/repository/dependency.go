package repository

import (
	"github.com/pujidjayanto/goginboilerplate/internal/repository/product"
	"github.com/pujidjayanto/goginboilerplate/internal/repository/purchase"
	"github.com/pujidjayanto/goginboilerplate/internal/repository/user"
	"github.com/pujidjayanto/goginboilerplate/pkg/db"
)

type Dependency struct {
	UserRepository     user.Repository
	ProductRepository  product.Repository
	PurchaseRepository purchase.Repository
}

// todo: should i pass pointer to db to avoid copy?
func NewDependency(dbHandler db.DatabaseHandler) Dependency {
	return Dependency{
		UserRepository:     user.NewRepository(dbHandler),
		ProductRepository:  product.NewRepository(dbHandler),
		PurchaseRepository: purchase.NewRepository(dbHandler),
	}
}
