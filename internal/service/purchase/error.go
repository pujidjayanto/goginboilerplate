package purchase

import "errors"

var (
	ErrProductNotFound     = errors.New("product not found")
	ErrInsufficientProduct = errors.New("product has 0 stock")
)
