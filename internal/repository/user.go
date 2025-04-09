package repository

import (
	"context"

	"github.com/t-lunch/t-lunch-backend/internal/models"
	"github.com/t-lunch/t-lunch-backend/pkg/db/tipo"
)

type UserRepository struct {
	tipodb *tipo.Users
}

func NewUserRepository(tdb *tipo.Users) *UserRepository {
	return &UserRepository{tipodb: tdb}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	return r.tipodb.AddUser(user.Name, user.Surname, user.Tg, user.Office, user.Login, user.Password, user.Emoji), nil
}

func (r *UserRepository) GetUser(ctx context.Context, id int64) (*models.User, error) {
	return r.tipodb.GetUser(id)
}

func (r *UserRepository) ListUsers(ctx context.Context) ([]*models.User, error) {
	return r.tipodb.ListUsers(), nil
}
