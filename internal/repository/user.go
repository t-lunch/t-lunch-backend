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

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) (int, error) {
	return r.tipodb.AddUser(user.Name, user.Surname, user.Tg, user.Office, user.Login, user.Password), nil
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (*models.User, error) {
	return r.tipodb.GetUser(id)
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	return r.tipodb.UpdateUser(user.ID, user.Name, user.Surname, user.Tg, user.Office, user.Login, user.Password)
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) (bool, error) {
	return r.tipodb.DeleteUser(id)
}

func (r *UserRepository) ListUsers(ctx context.Context) ([]*models.User, error) {
	return r.tipodb.ListUsers(), nil
}
