package service

import (
	"github.com/pujidjayanto/goginboilerplate/internal/repository"
	"github.com/pujidjayanto/goginboilerplate/internal/service/product"
	"github.com/pujidjayanto/goginboilerplate/internal/service/user"
)

type Dependency struct {
	UserService    user.Service
	ProductService product.Service
}

func NewDependency(repositories repository.Dependency) Dependency {
	return Dependency{
		UserService:    user.NewService(repositories.UserRepository),
		ProductService: product.NewService(repositories.ProductRepository),
	}
}
