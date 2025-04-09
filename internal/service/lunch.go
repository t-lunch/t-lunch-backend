package service

import (
	"context"

	"github.com/t-lunch/t-lunch-backend/internal/models"
)

type LunchRepo interface {
	CreateLunch(ctx context.Context, lunch *models.Lunch) (int, error)
	GetLunch(ctx context.Context, id int) (*models.Lunch, error)
	UpdateLunch(ctx context.Context, lunch *models.Lunch) (*models.Lunch, error)
	DeleteLunch(ctx context.Context, id int) (bool, error)
	ListLunches(ctx context.Context) ([]*models.Lunch, error)
}

type LunchService struct {
	Repo LunchRepo
}

func NewLunchService(repo LunchRepo) *LunchService {
	return &LunchService{
		Repo: repo,
	}
}

func (s *LunchService) CreateLunch(ctx context.Context, lunch *models.Lunch) (int, error) {
	return s.Repo.CreateLunch(ctx, lunch)
}

func (s *LunchService) GetLunch(ctx context.Context, id int) (*models.Lunch, error) {
	return s.Repo.GetLunch(ctx, id)
}

func (s *LunchService) UpdateLunch(ctx context.Context, lunch *models.Lunch) (*models.Lunch, error) {
	return s.Repo.UpdateLunch(ctx, lunch)
}

func (s *LunchService) DeleteLunch(ctx context.Context, id int) (bool, error) {
	return s.Repo.DeleteLunch(ctx, id)
}
func (s *LunchService) ListLunches(ctx context.Context) ([]*models.Lunch, error) {
	return s.Repo.ListLunches(ctx)
}
