package service

import (
	"context"
	"errors"

	"github.com/t-lunch/t-lunch-backend/internal/models"
)

type AuthRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (int64, error)
	CheckPassword(ctx context.Context, email, password string) (bool, error)
	GenerateAccessToken(ctx context.Context) string
	GenerateRefreshToken(ctx context.Context) string
	ValidateRefreshToken(ctx context.Context, token string) (bool, error)
}

type AuthService struct {
	repo AuthRepo
}

func NewAuthService(repo AuthRepo) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Registration(ctx context.Context, user *models.User) (*models.User, error) {
	// TODO проверка, что пользователя с таким логином еще нет
	if user.Name == "" || user.Surname == "" || user.Tg == "" || user.Office == "" || user.Emoji == "" || user.Email == "" || user.Password == "" {
		return nil, errors.New("Все поля обязательны")
	}

	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.New("error AuthRepo: CreateUser")
	}

	return user, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, string, error) {
	if email == "" || password == "" {
		return "", "", errors.New("Все поля обязательны")
	}

	_, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", errors.New("error AuthRepo: GetUserByEmail")
	}

	isCorrect, err := s.repo.CheckPassword(ctx, email, password)
	if err != nil {
		return "", "", errors.New("error AuthRepo: CheckPassword")
	}
	if !isCorrect {
		return "", "", errors.New("Неверные данные")
	}

	return s.repo.GenerateAccessToken(ctx), s.repo.GenerateRefreshToken(ctx), nil
}

func (s *AuthService) Refresh(ctx context.Context, token string) (string, error) {
	if token == "" {
		return "", errors.New("Все поля обязательны")
	}

	isCorrect, err := s.repo.ValidateRefreshToken(ctx, token)
	if err != nil {
		return "", errors.New("error AuthRepo: ValidateRefreshToken")
	}
	if !isCorrect {
		return "", errors.New("Неверный refresh-токен")
	}

	return s.repo.GenerateAccessToken(ctx), nil
}
