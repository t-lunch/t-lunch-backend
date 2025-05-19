package service

import (
	"context"
	"slices"

	"github.com/t-lunch/t-lunch-backend/internal/errors"
	"github.com/t-lunch/t-lunch-backend/internal/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserPasswordByEmail(ctx context.Context, email string) (int64, string, error)
	IsUserWithEmailExist(ctx context.Context, email string) (bool, error)
	GetUsersByIDs(ctx context.Context, ids []int64) ([]*models.UserResponse, error)
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
