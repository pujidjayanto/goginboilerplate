package purchase

import (
	"time"

	"github.com/pujidjayanto/goginboilerplate/internal/repository/product"
	"github.com/pujidjayanto/goginboilerplate/internal/repository/user"
)

type Purchase struct {
	ID           uint            `gorm:"primaryKey"`
	UserID       uint            `gorm:"not null"`
	ProductID    uint            `gorm:"not null"`
	PurchaseDate time.Time       `gorm:"autoCreateTime"`
	Quantity     int             `gorm:"not null"`
	CreatedAt    time.Time       `gorm:"autoCreateTime"`
	User         user.User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Product      product.Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}
