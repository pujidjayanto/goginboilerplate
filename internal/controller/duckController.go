package controller

import (
	"context"

	"github.com/pujidjayanto/goginboilerplate/internal/service/duck"
)

type DuckController interface {
	Create(context.Context) error
}

type duckController struct {
	service duck.Service
}

func NewDuckController(service duck.Service) DuckController {
	return &duckController{service: service}
}

func (c *duckController) Create(ctx context.Context) error {
	return nil
}
