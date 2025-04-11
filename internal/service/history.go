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

func (s *HistoryService) RateLunch(ctx context.Context, userID, lunchID int64, isLiked bool) error {
	history := &models.History{
		UserID:  userID,
		LunchID: lunchID,
		IsLiked: isLiked,
	}

	// Проверяем, существует ли уже запись
	histories, err := s.historyRepo.GetHistoryByUser(ctx, userID)
	if err != nil {
		return err
	}

	// Если запись уже есть, обновляем её
	for _, h := range histories {
		if h.LunchID == lunchID {
			history.ID = h.ID
			return s.historyRepo.UpdateHistory(ctx, history)
		}
	}

	// Если записи нет, создаём новую
	return s.historyRepo.CreateHistory(ctx, history)
}

func (s *HistoryService) GetHistoryByUser(ctx context.Context, userID int64) ([]models.History, error) {
	return s.historyRepo.GetHistoryByUser(ctx, userID)
}
