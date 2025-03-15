package purchase

import (
	"time"

	"github.com/pujidjayanto/goginboilerplate/internal/repository/product"
	"github.com/pujidjayanto/goginboilerplate/internal/repository/user"
)

type Purchase struct {
	ID           uint            `gorm:"primaryKey"`
	UserId       uint            `gorm:"not null"`
	ProductId    uint            `gorm:"not null"`
	PurchaseDate time.Time       `gorm:"autoCreateTime"` // todo: need to investigate on this
	Quantity     int             `gorm:"not null"`
	CreatedAt    time.Time       `gorm:"autoCreateTime"`
	User         user.User       `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	Product      product.Product `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE"`
}
