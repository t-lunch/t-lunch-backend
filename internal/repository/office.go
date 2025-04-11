package repository

import (
	"context"

	"github.com/t-lunch-backend/internal/models"

	"database/sql"
)

type PostgresOfficeRepository struct {
	db *sql.DB
}

func NewOfficeRepository(db *sql.DB) OfficeRepository {
	return &PostgresOfficeRepository{db: db}
}

func (r *PostgresOfficeRepository) CreateOffice(ctx context.Context, office *models.Office) error {
	query := "INSERT INTO offices (address) VALUES ($1) RETURNING id"
	return r.db.QueryRowContext(ctx, query, office.Address).Scan(&office.ID)
}

func (r *PostgresOfficeRepository) GetOfficeByID(ctx context.Context, id int64) (*models.Office, error) {
	var office models.Office
	query := "SELECT id, address FROM offices WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&office.ID, &office.Address)
	if err != nil {
		return nil, err
	}
	return &office, nil
}
