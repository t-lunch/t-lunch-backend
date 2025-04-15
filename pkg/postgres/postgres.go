package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewDB(ctx context.Context, connectionString string) (*DB, error) {
	pollCfg, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	poll, err := pgxpool.NewWithConfig(ctx, pollCfg)
	if err != nil {
		return nil, err
	}

	return &DB{Pool: poll}, nil
}
