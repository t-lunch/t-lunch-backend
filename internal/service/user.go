package service

import (
	"context"
	"slices"

	"github.com/t-lunch/t-lunch-backend/internal/errors"
	"github.com/t-lunch/t-lunch-backend/internal/models"
)

//go:generate mockgen -destination=./mocks/user.go -package=mocks . UserRepo

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserPasswordByEmail(ctx context.Context, email string) (int64, string, error)
	IsUserWithEmailExist(ctx context.Context, email string) (bool, error)
	GetUsersByIDs(ctx context.Context, ids []int64) ([]*models.UserResponse, error)
	GetUserByID(ctx context.Context, id int64) (*models.UserResponse, error)
	UpdateUserByID(ctx context.Context, id int64, updates map[string]interface{}) error
}

type UserService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUsersByIDs(ctx context.Context, userIDs []int64) ([]*models.UserResponse, error) {
	if len(userIDs) == 0 || slices.ContainsFunc(userIDs, func(userId int64) bool {
		return userId <= 0
	}) {
		return nil, errors.ErrInvalidRequest
	}

	users, err := s.userRepo.GetUsersByIDs(ctx, userIDs)
	if err != nil {
		return nil, errors.NewErrRepository("userRepo", "GetUsersByIDs", err)
	}

	return users, nil
}

func (s *UserService) GetUserByID(ctx context.Context, userID int64) (*models.UserResponse, error) {
	if userID <= 0 {
		return nil, errors.ErrInvalidRequest
	}

	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, errors.NewErrRepository("userRepo", "GetUserByID", err)
	}

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.UserResponse) (*models.UserResponse, error) {
	if user.ID <= 0 {
		return nil, errors.ErrInvalidRequest
	}

	err := s.userRepo.UpdateUserByID(ctx, user.ID, map[string]interface{}{
		"name":    user.Name,
		"surname": user.Surname,
		"tg":      user.Tg,
		"office":  user.Office,
		"emoji":   user.Emoji,
	})
	if err != nil {
		return nil, errors.NewErrRepository("userRepo", "UpdateUserByID", err)
	}

	updatedUser, err := s.userRepo.GetUserByID(ctx, user.ID)
	if err != nil {
		return nil, errors.NewErrRepository("userRepo", "GetUserByID", err)
	}

	return updatedUser, nil
}
