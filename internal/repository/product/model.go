package product

import (
	"time"

	"github.com/pujidjayanto/goginboilerplate/pkg/jsonb"
)

type Product struct {
	ID             uint       `gorm:"primaryKey"`
	Name           string     `gorm:"not null"`
	Price          float64    `gorm:"not null"`
	Quantity       int        `gorm:"not null"`
	ProductDetails jsonb.JSON `gorm:"type:jsonb"`
	CreatedAt      time.Time  `gorm:"autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime"`
}
