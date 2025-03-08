package product

import "time"

type Product struct {
	ID             uint           `gorm:"primaryKey"`
	Name           string         `gorm:"not null"`
	Price          float64        `gorm:"not null"`
	Quantity       int            `gorm:"not null"`
	ProductDetails map[string]any `gorm:"type:jsonb"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
}
