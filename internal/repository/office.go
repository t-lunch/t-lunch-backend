package repository

import (
	"context"

	"github.com/t-lunch/t-lunch-backend/internal/models"
	"github.com/t-lunch/t-lunch-backend/pkg/db/tipo"
)

type OfficeRepository struct {
	tipodb *tipo.Offices
}

func NewOfficeRepository(tdb *tipo.Offices) *OfficeRepository {
	return &OfficeRepository{tipodb: tdb}
}

func (r *OfficeRepository) CreateOffice(ctx context.Context, office *models.Office) (int, error) {
	return r.tipodb.AddOffice(office.Address), nil
}

func (r *OfficeRepository) GetOffice(ctx context.Context, id int) (*models.Office, error) {
	return r.tipodb.GetOffice(id)
}

func (r *OfficeRepository) ListOffices(ctx context.Context) ([]*models.Office, error) {
	return r.tipodb.ListOffices(), nil
}
