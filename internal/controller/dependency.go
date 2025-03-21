package controller

import (
	"github.com/pujidjayanto/goginboilerplate/internal/controller/product"
	"github.com/pujidjayanto/goginboilerplate/internal/controller/purchase"
	"github.com/pujidjayanto/goginboilerplate/internal/controller/user"
	"github.com/pujidjayanto/goginboilerplate/internal/service"
)

type Dependency struct {
	UserController     user.Controller
	ProductController  product.Controller
	PurchaseController purchase.Controller
}

func NewDependency(services service.Dependency) Dependency {
	return Dependency{
		UserController:     user.NewController(services.UserService),
		ProductController:  product.NewController(services.ProductService),
		PurchaseController: purchase.NewController(services.PurchaseService),
	}
}
