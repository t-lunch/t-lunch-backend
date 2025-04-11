package service

import (
	"context"

	"github.com/t-lunch-backend/pkg/models"
	"github.com/t-lunch-backend/pkg/repository"
)

type OfficeService struct {
	officeRepo repository.OfficeRepository
}

func NewOfficeService(officeRepo repository.OfficeRepository) *OfficeService {
	return &OfficeService{officeRepo: officeRepo}
}

func (s *OfficeService) CreateOffice(ctx context.Context, office *models.Office) error {
	return s.officeRepo.CreateOffice(ctx, office)
}

func (s *OfficeService) GetOfficeByID(ctx context.Context, id int64) (*models.Office, error) {
	return s.officeRepo.GetOfficeByID(ctx, id)
}
