package service

import (
	"context"

	"github.com/t-lunch-backend/pkg/models"
	"github.com/t-lunch-backend/pkg/repository"
)

type TlunchService struct {
	lunchRepo repository.LunchRepository
}

func NewTlunchService(lunchRepo repository.LunchRepository) *TlunchService {
	return &TlunchService{lunchRepo: lunchRepo}
}

func (s *TlunchService) CreateLunch(ctx context.Context, lunch *models.Lunch) error {
	return s.lunchRepo.CreateLunch(ctx, lunch)
}

func (s *TlunchService) GetLunches(ctx context.Context, userID int64, offset, limit int) ([]models.Lunch, error) {
	return s.lunchRepo.GetLunches(ctx, userID, offset, limit)
}

func (s *TlunchService) GetLunchByID(ctx context.Context, lunchID int64) (*models.Lunch, error) {
	return s.lunchRepo.GetLunchByID(ctx, lunchID)
}

func (s *TlunchService) JoinLunch(ctx context.Context, lunchID, userID int64) error {
	return s.lunchRepo.JoinLunch(ctx, lunchID, userID)
}

func (s *TlunchService) LeaveLunch(ctx context.Context, lunchID, userID int64) error {
	return s.lunchRepo.LeaveLunch(ctx, lunchID, userID)
}

func (s *TlunchService) RateLunch(ctx context.Context, lunchID, userID int64, rating float64) error {
	return s.lunchRepo.RateLunch(ctx, lunchID, userID, rating)
}

func (s *TlunchService) GetLunchHistory(ctx context.Context, lunchID int64) (float64, error) {
	return s.lunchRepo.GetLunchHistory(ctx, lunchID)
}
