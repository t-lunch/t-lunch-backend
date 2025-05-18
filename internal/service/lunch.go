package service

import (
	"context"
	goerrors "errors"
	"time"

	"github.com/t-lunch/t-lunch-backend/internal/errors"
	"github.com/t-lunch/t-lunch-backend/internal/models"
	"gorm.io/gorm"
)

type LunchRepo interface {
	CreateLunch(ctx context.Context, lunch *models.Lunch) error
	GetLunches(ctx context.Context, userID int64, offset, limit int) ([]*models.Lunch, error)
	GetLunchByID(ctx context.Context, id int64) (*models.Lunch, error)
	GetLunchIdByUserID(ctx context.Context, userID int64) (int64, error)
}

type LunchService struct {
	lunchRepo LunchRepo
}

func NewLunchService(lunchRepo LunchRepo) *LunchService {
	return &LunchService{lunchRepo: lunchRepo}
}

func (s *LunchService) CreateLunch(ctx context.Context, userID int64, place string, lunchTime time.Time, description string) (*models.Lunch, error) {
	if userID <= 0 || place == "" {
		return nil, errors.ErrInvalidRequest
	}
	// if !ValidTime(ctx, time.Now(), lunchTime) {
	// 	return nil, errors.New("invalid lunch time")
	// }

	lunch := &models.Lunch{
		CreatorID:            userID,
		Place:                place,
		Time:                 lunchTime,
		NumberOfParticipants: 1,
		Description:          description,
		Participants:         []int64{userID},
	}

	err := s.lunchRepo.CreateLunch(ctx, lunch)
	if err != nil {
		return nil, errors.NewErrRepository("lunchRepo", "CreateLunch", err)
	}

	createdLunch, err := s.lunchRepo.GetLunchByID(ctx, lunch.ID)
	if err != nil {
		return nil, errors.NewErrRepository("lunchRepo", "GetLunchByID", err)
	}

	return createdLunch, nil
}

func (s *LunchService) GetLunches(ctx context.Context, userID int64, offset, limit int) ([]*models.Lunch, int64, error) {
	if userID <= 0 {
		return nil, 0, errors.ErrInvalidRequest
	}

	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 5
	}
	if limit > 10 {
		limit = 10
	}

	lunchID, err := s.lunchRepo.GetLunchIdByUserID(ctx, userID)
	if err != nil && !goerrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, errors.NewErrRepository("lunchRepo", "GetLunchIdByUserID", err)
	}

	lunches, err := s.lunchRepo.GetLunches(ctx, userID, offset, limit)
	if err != nil {
		return nil, 0, errors.NewErrRepository("lunchRepo", "GetLunches", err)
	}

	return lunches, lunchID, nil
}

func ValidTime(ctx context.Context, now, lunchTime time.Time) bool {
	beginLunchTime := time.Date(now.Year(), now.Month(), now.Day(), 11, 0, 0, 0, now.Location())
	endLunchTime := time.Date(now.Year(), now.Month(), now.Day(), 14, 0, 0, 0, now.Location())
	return !lunchTime.Before(beginLunchTime) && !lunchTime.After(endLunchTime)
}

// func (s *TlunchService) GetLunchByID(ctx context.Context, lunchID int64) (*models.Lunch, error) {
// 	return s.lunchRepo.GetLunchByID(ctx, lunchID)
// }

// func (s *TlunchService) JoinLunch(ctx context.Context, lunchID, userID int64) error {
// 	return s.lunchRepo.JoinLunch(ctx, lunchID, userID)
// }

// func (s *TlunchService) LeaveLunch(ctx context.Context, lunchID, userID int64) error {
// 	return s.lunchRepo.LeaveLunch(ctx, lunchID, userID)
// }

// func (s *TlunchService) GetLunchHistory(ctx context.Context, lunchID int64) (float64, error) {
// 	return s.lunchRepo.GetLunchHistory(ctx, lunchID)
// }
