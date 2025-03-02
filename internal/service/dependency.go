package service

import (
	"github.com/pujidjayanto/goginboilerplate/internal/repository"
	"github.com/pujidjayanto/goginboilerplate/internal/service/duck"
)

type Dependency struct {
	DuckService duck.Service
}

func NewDependency(repositories repository.Dependency) Dependency {
	return Dependency{
		DuckService: duck.NewService(repositories.DuckRepository),
	}
}
