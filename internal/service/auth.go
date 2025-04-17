package service

import (
	"context"
	"errors"

	"github.com/t-lunch/t-lunch-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserPasswordByEmail(ctx context.Context, email string) (string, error)
}

type AuthRepo interface {
	GenerateAccessToken(ctx context.Context) string
	GenerateRefreshToken(ctx context.Context) string
}

type AuthService struct {
	userRepo UserRepo
	authRepo AuthRepo
}

func NewAuthService(userRepo UserRepo, authRepo AuthRepo) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		authRepo: authRepo,
	}
}

func (s *AuthService) Registration(ctx context.Context, user *models.User) (*models.User, error) {
	// TODO проверка, что пользователя с таким логином еще нет
	if user.Name == "" || user.Surname == "" || user.Tg == "" || user.Office == "" || user.Emoji == "" || user.Email == "" || user.HashedPassword == "" {
		return nil, errors.New("все поля обязательны")
	}

	if _, err := s.userRepo.GetUserPasswordByEmail(ctx, user.Email); err == nil {
		return nil, errors.New("пользователь с таким email уже существует")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.HashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("ошибка хэш пароля")
	}
	user.HashedPassword = string(hashedPassword)

	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.New("error userRepo: CreateUser")
	}

	return user, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, string, error) {
	if email == "" || password == "" {
		return "", "", errors.New("все поля обязательны")
	}

	hashedPassword, err := s.userRepo.GetUserPasswordByEmail(ctx, email)
	if err != nil {
		return "", "", errors.New("error userRepo: GetUserPasswordByEmail")
	}

	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return "", "", errors.New("error: CompareHashAndPassword")
	}

	return s.authRepo.GenerateAccessToken(ctx), s.authRepo.GenerateRefreshToken(ctx), nil
}

func (s *AuthService) Refresh(ctx context.Context, token string) (string, error) {
	if token == "" {
		return "", errors.New("все поля обязательны")
	}

	// TODO
	if token != "mock_refresh_token" { // Замените на реализацию JWT
		return "", errors.New("неверный refresh-токен")
	}

	return s.authRepo.GenerateAccessToken(ctx), nil
}
