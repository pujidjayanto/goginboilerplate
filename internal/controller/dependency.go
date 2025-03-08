package controller

import (
	"github.com/pujidjayanto/goginboilerplate/internal/controller/user"
	"github.com/pujidjayanto/goginboilerplate/internal/service"
)

type Dependency struct {
	UserController user.Controller
}

func NewDependency(services service.Dependency) Dependency {
	return Dependency{
		UserController: user.NewController(services.UserService),
	}
}
