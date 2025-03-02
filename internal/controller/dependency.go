package controller

import "github.com/pujidjayanto/goginboilerplate/internal/service"

type Dependency struct {
	DuckController DuckController
}

func NewDependency(services service.Dependency) Dependency {
	return Dependency{
		DuckController: NewDuckController(services.DuckService),
	}
}
