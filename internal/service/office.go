package service

import (
	"context"

	"github.com/t-lunch/t-lunch-backend/internal/models"
)

type OfficeRepo interface {
	CreateOffice(ctx context.Context, office *models.Office) (int, error)
	GetOffice(ctx context.Context, id int) (*models.Office, error)
	UpdateOffice(ctx context.Context, office *models.Office) (*models.Office, error)
	DeleteOffice(ctx context.Context, id int) (bool, error)
	ListOffices(ctx context.Context) ([]*models.Office, error)
}

type OfficeService struct {
	Repo OfficeRepo
}

func NewOfficeService(repo OfficeRepo) *OfficeService {
	return &OfficeService{
		Repo: repo,
	}
}

func (s *OfficeService) CreateOffice(ctx context.Context, office *models.Office) (int, error) {
	return s.Repo.CreateOffice(ctx, office)
}

func (s *OfficeService) GetOffice(ctx context.Context, id int) (*models.Office, error) {
	return s.Repo.GetOffice(ctx, id)
}

func (s *OfficeService) UpdateOffice(ctx context.Context, office *models.Office) (*models.Office, error) {
	return s.Repo.UpdateOffice(ctx, office)
}

func (s *OfficeService) DeleteOffice(ctx context.Context, id int) (bool, error) {
	return s.Repo.DeleteOffice(ctx, id)
}

func (s *OfficeService) ListOffices(ctx context.Context) ([]*models.Office, error) {
	return s.Repo.ListOffices(ctx)
}
