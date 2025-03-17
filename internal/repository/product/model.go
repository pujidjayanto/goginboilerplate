package product

import (
	"time"

	"github.com/pujidjayanto/goginboilerplate/pkg/jsonb"
	"github.com/shopspring/decimal"
)

type Product struct {
	Id             uint            `gorm:"primaryKey"`
	Name           string          `gorm:"not null"`
	Price          decimal.Decimal `gorm:"not null type:decimal(10,2)"`
	Quantity       int             `gorm:"not null"`
	ProductDetails jsonb.JSON      `gorm:"type:jsonb"`
	CreatedAt      time.Time       `gorm:"autoCreateTime"`
	UpdatedAt      time.Time       `gorm:"autoUpdateTime"`
}

type ProductFilter struct {
	ProductName string
}
