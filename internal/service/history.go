package service

import (
	"context"

	"github.com/t-lunch-backend/pkg/models"
	"github.com/t-lunch-backend/pkg/repository"
)

type HistoryService struct {
	historyRepo repository.HistoryRepository
}

func NewHistoryService(historyRepo repository.HistoryRepository) *HistoryService {
	return &HistoryService{historyRepo: historyRepo}
}

func (s *HistoryService) CreateHistory(ctx context.Context, history *models.History) error {
	return s.historyRepo.CreateHistory(ctx, history)
}

func (s *HistoryService) GetHistoryByUser(ctx context.Context, userID int64) ([]models.History, error) {
	return s.historyRepo.GetHistoryByUser(ctx, userID)
}

func (s *HistoryService) UpdateHistory(ctx context.Context, history *models.History) error {
	return s.historyRepo.UpdateHistory(ctx, history)
}
