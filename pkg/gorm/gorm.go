package gorm

import (
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(ctx context.Context, connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
