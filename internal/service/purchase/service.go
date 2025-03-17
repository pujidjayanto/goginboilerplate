package purchase

import (
	"context"
	"errors"
	"fmt"

	"github.com/pujidjayanto/goginboilerplate/internal/dto"
	"github.com/pujidjayanto/goginboilerplate/internal/repository/product"
	"github.com/pujidjayanto/goginboilerplate/internal/repository/purchase"
	"github.com/pujidjayanto/goginboilerplate/pkg/db"
	"github.com/pujidjayanto/goginboilerplate/pkg/logger"
	"gorm.io/gorm"
)

type Service interface {
	MakePurchase(context.Context, dto.CreatePurchaseRequest) error
}

type service struct {
	purchaseRepository purchase.Repository
	productRepository  product.Repository
	dbHandler          db.DatabaseHandler
}

func (s *service) MakePurchase(ctx context.Context, req dto.CreatePurchaseRequest) error {
	product, err := s.productRepository.FindById(ctx, req.ProductId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrProductNotFound
		}
		return fmt.Errorf("failed to get product, %v", err)
	}

	if product.Quantity == 0 || req.Quantity > product.Quantity {
		return ErrInsufficientProduct
	}

	// example usage of transaction
	err = s.dbHandler.RunTransaction(ctx, func(innerCtx context.Context) error {
		newPurchase := purchase.Purchase{
			UserId:    req.UserId,
			ProductId: req.ProductId,
			Quantity:  req.Quantity,
		}
		err := s.purchaseRepository.Create(innerCtx, newPurchase)
		if err != nil {
			return fmt.Errorf("failed to create purchase, %v", err)
		}

		currentStock := product.Quantity - req.Quantity
		product.Quantity = currentStock

		err = s.productRepository.Update(innerCtx, product)
		if err != nil {
			return fmt.Errorf("failed to update product, %v", err)
		}

		// example usage of logger
		logger.Info("purchase has created", "purchaseId", newPurchase.Id)
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to create purchase, %v", err)
	}

	return nil
}

func NewService(
	purr purchase.Repository,
	pur product.Repository,
	dbh db.DatabaseHandler,
) Service {
	return &service{
		purchaseRepository: purr,
		productRepository:  pur,
		dbHandler:          dbh,
	}
}
