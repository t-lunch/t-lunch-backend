package service

import (
	"context"

	"github.com/t-lunch/t-lunch-backend/internal/models"
)

type HistoryRepo interface {
	CreateHistory(ctx context.Context, history *models.History) (int, error)
	GetHistory(ctx context.Context, id int) (*models.History, error)
	UpdateHistory(ctx context.Context, history *models.History) (*models.History, error)
	DeleteHistory(ctx context.Context, id int) (bool, error)
	ListHistory(ctx context.Context) ([]*models.History, error)
}

type HistoryService struct {
	Repo HistoryRepo
}

func NewHistoryService(repo HistoryRepo) *HistoryService {
	return &HistoryService{
		Repo: repo,
	}
}

func (s *HistoryService) CreateHistory(ctx context.Context, history *models.History) (int, error) {
	return s.Repo.CreateHistory(ctx, history)
}

func (s *HistoryService) GetHistory(ctx context.Context, id int) (*models.History, error) {
	return s.Repo.GetHistory(ctx, id)
}

func (s *HistoryService) UpdateHistory(ctx context.Context, history *models.History) (*models.History, error) {
	return s.Repo.UpdateHistory(ctx, history)
}

func (s *HistoryService) DeleteHistory(ctx context.Context, id int) (bool, error) {
	return s.Repo.DeleteHistory(ctx, id)
}
func (s *HistoryService) ListHistory(ctx context.Context) ([]*models.History, error) {
	return s.Repo.ListHistory(ctx)
}
