package service

import (
	"github.com/pujidjayanto/goginboilerplate/internal/repository"
	"github.com/pujidjayanto/goginboilerplate/internal/service/user"
)

type Dependency struct {
	UserService user.Service
}

func NewDependency(repositories repository.Dependency) Dependency {
	return Dependency{
		UserService: user.NewService(repositories.UserRepository),
	}
}
