package product

import (
	"time"

	"github.com/pujidjayanto/goginboilerplate/pkg/jsonb"
)

type Product struct {
	ID             uint       `gorm:"primaryKey"`
	Name           string     `gorm:"not null"`
	Price          float64    `gorm:"not null"` // todo: need to use custom type for price
	Quantity       int        `gorm:"not null"`
	ProductDetails jsonb.JSON `gorm:"type:jsonb"`
	CreatedAt      time.Time  `gorm:"autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime"`
}

type ProductFilter struct {
	ProductName string
}
