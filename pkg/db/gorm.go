package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseHandler interface {
	GetDB(ctx context.Context) *gorm.DB
	RunTransaction(ctx context.Context, fc func(ctx context.Context) error) error
	Ping(ctx context.Context) error
	Close() error
}

type databaseHandler struct {
	db  *gorm.DB
	sql *sql.DB
}

type contextKey string

const txKey = contextKey("DBTX")

func InitDatabaseHandler(dsn string) (DatabaseHandler, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc:                func() time.Time { return time.Now().UTC() },
	})
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	dbPool, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error opening database pool: %w", err)
	}

	// Set connection pool settings
	dbPool.SetMaxIdleConns(10)
	dbPool.SetMaxOpenConns(100)
	dbPool.SetConnMaxLifetime(time.Hour)

	// Database ping
	err = dbPool.Ping()
	if err != nil {
		return nil, fmt.Errorf("can't ping the db, WTF %v", err)
	}

	return &databaseHandler{db: db, sql: dbPool}, nil
}

func (h *databaseHandler) GetDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey).(*gorm.DB); ok {
		return tx.WithContext(ctx)
	}
	return h.db.WithContext(ctx)
}

func (h *databaseHandler) Ping(ctx context.Context) error {
	return h.sql.PingContext(ctx)
}

func (h *databaseHandler) RunTransaction(ctx context.Context, fc func(ctx context.Context) error) error {
	return h.db.Transaction(func(tx *gorm.DB) error {
		// Add the transaction to the context
		return fc(context.WithValue(ctx, txKey, tx))
	})
}

func (h *databaseHandler) Close() error {
	err := h.sql.Close()
	if err != nil {
		return err
	}

	return nil
}
