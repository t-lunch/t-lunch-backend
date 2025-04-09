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

func (r *HistoryRepository) CreateHistory(ctx context.Context, history *models.History) (int, error) {
	return r.tipodb.AddHistory(history.UserID, history.LunchID, history.Date), nil
}

func (r *HistoryRepository) GetHistory(ctx context.Context, id int) (*models.History, error) {
	return r.tipodb.GetHistory(id)
}

func (r *HistoryRepository) UpdateHistory(ctx context.Context, history *models.History) (*models.History, error) {
	return r.tipodb.UpdateHistory(history.ID, history.UserID, history.LunchID, history.Date)
}

func (r *HistoryRepository) DeleteHistory(ctx context.Context, id int) (bool, error) {
	return r.tipodb.DeleteHistory(id)
}

func (r *HistoryRepository) ListHistory(ctx context.Context) ([]*models.History, error) {
	return r.tipodb.ListHistory(), nil
}
