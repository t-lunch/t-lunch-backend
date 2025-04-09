package repository

import (
	"context"

	"github.com/t-lunch/t-lunch-backend/internal/models"
	"github.com/t-lunch/t-lunch-backend/pkg/db/tipo"
)

type LunchRepository struct {
	tipodb *tipo.Lunches
}

func NewLunchRepository(tdb *tipo.Lunches) *LunchRepository {
	return &LunchRepository{tipodb: tdb}
}

func (r *LunchRepository) CreateLunch(ctx context.Context, lunch *models.Lunch) (int64, error) {
	return r.tipodb.AddLunch(lunch.Creator, lunch.Time, lunch.Place, lunch.Optional, lunch.Participants), nil
}

func (r *LunchRepository) GetLunch(ctx context.Context, id int64) (*models.Lunch, error) {
	return r.tipodb.GetLunch(id)
}

func (r *LunchRepository) UpdateLunch(ctx context.Context, lunch *models.Lunch) (*models.Lunch, error) {
	return r.tipodb.UpdateLunch(lunch.ID, lunch.Creator, lunch.Time, lunch.Place, lunch.Optional, lunch.Participants)
}

func (r *LunchRepository) ListLunches(ctx context.Context) ([]*models.Lunch, error) {
	return r.tipodb.ListLunches(), nil
}
