package service

import (
	"github.com/pujidjayanto/goginboilerplate/internal/repository"
	"github.com/pujidjayanto/goginboilerplate/internal/service/product"
	"github.com/pujidjayanto/goginboilerplate/internal/service/purchase"
	"github.com/pujidjayanto/goginboilerplate/internal/service/user"
	"github.com/pujidjayanto/goginboilerplate/pkg/db"
)

type Dependency struct {
	UserService     user.Service
	ProductService  product.Service
	PurchaseService purchase.Service
}

func NewDependency(repositories repository.Dependency, dbHandler db.DatabaseHandler) Dependency {
	return Dependency{
		UserService:     user.NewService(repositories.UserRepository),
		ProductService:  product.NewService(repositories.ProductRepository),
		PurchaseService: purchase.NewService(repositories.PurchaseRepository, repositories.ProductRepository, dbHandler),
	}
}
