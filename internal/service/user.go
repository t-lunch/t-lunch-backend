package service

import (
	"context"

	"github.com/t-lunch/t-lunch-backend/internal/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) (int, error)
	GetUser(ctx context.Context, id int) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id int) (bool, error)
	ListUsers(ctx context.Context) ([]*models.User, error)
}

type UserService struct {
	Repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) (int, error) {
	return s.Repo.CreateUser(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id int) (*models.User, error) {
	return s.Repo.GetUser(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	return s.Repo.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int) (bool, error) {
	return s.Repo.DeleteUser(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context) ([]*models.User, error) {
	return s.Repo.ListUsers(ctx)
}
