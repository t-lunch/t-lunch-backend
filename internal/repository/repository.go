package repository

import (
	"context"

	"github.com/t-lunch-backend/internal/models"
)

type LunchRepository interface {
	CreateLunch(ctx context.Context, lunch *models.Lunch) error
	GetLunches(ctx context.Context, userID int64, offset, limit int) ([]models.Lunch, error)
	GetLunchByID(ctx context.Context, lunchID int64) (*models.Lunch, error)
	JoinLunch(ctx context.Context, lunchID, userID int64) error
	LeaveLunch(ctx context.Context, lunchID, userID int64) error
	RateLunch(ctx context.Context, userID, lunchID int64, isLiked bool) error
	GetLunchHistory(ctx context.Context, userID int64) ([]models.LunchFeedback, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, userID int64) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
}
