package repository

import (
	"context"

	"github.com/t-lunch/t-lunch-backend/internal/models"
	"github.com/t-lunch/t-lunch-backend/pkg/db/tipo"
)

type HistoryRepository struct {
	tipodb *tipo.Histories
}

func NewHistoryRepository(tdb *tipo.Histories) *HistoryRepository {
	return &HistoryRepository{tipodb: tdb}
}

func (r *HistoryRepository) CreateHistory(ctx context.Context, history *models.History) (int64, error) {
	return r.tipodb.AddHistory(history.UserID, history.LunchID, history.IsLiked), nil
}

func (r *HistoryRepository) GetHistory(ctx context.Context, id int64) (*models.History, error) {
	return r.tipodb.GetHistory(id)
}

func (r *HistoryRepository) ListHistory(ctx context.Context) ([]*models.History, error) {
	return r.tipodb.ListHistory(), nil
}
