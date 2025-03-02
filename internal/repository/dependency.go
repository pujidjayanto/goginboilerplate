package repository

import (
	"github.com/pujidjayanto/goginboilerplate/internal/repository/duck"
	"github.com/pujidjayanto/goginboilerplate/pkg/db"
)

type Dependency struct {
	DuckRepository duck.Repository
}

func NewDependency(dbHandler db.DatabaseHandler) Dependency {
	return Dependency{
		DuckRepository: duck.NewRepository(dbHandler),
	}
}
