package user

import (
	"context"

	"github.com/pujidjayanto/goginboilerplate/internal/service/user"
)

type Controller interface {
	Create(context.Context) error
}

type controller struct {
	userService user.Service
}

func NewController(uss user.Service) Controller {
	return &controller{userService: uss}
}

func (c *controller) Create(ctx context.Context) error {
	return nil
}
