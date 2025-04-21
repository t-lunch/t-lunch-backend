package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/t-lunch/t-lunch-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserPasswordByEmail(ctx context.Context, email string) (string, error)
}

type AuthRepo interface {
	GenerateToken(ctx context.Context, secret string, expiration time.Duration, id int64) (string, error)
	ValidateToken(ctx context.Context, secret, token string) (int64, bool)
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

	accessToken, err := s.authRepo.GenerateToken(ctx, "secret", time.Hour*24, 1)
	if err != nil {
		return "", "", errors.New("error authRepo: GenerateToken access")
	}

	refreshToken, err := s.authRepo.GenerateToken(ctx, "secret", time.Hour*24*30, 1)
	if err != nil {
		return "", "", errors.New("error authRepo: GenerateToken refresh")
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) Refresh(ctx context.Context, token string) (string, error) {
	if token == "" {
		return "", errors.New("все поля обязательны")
	}

	id, ok := s.authRepo.ValidateToken(ctx, "secret", token)
	if !ok {
		return "", fmt.Errorf("error authRepo: ValidateToken %d", id)
	}

	accessToken, err := s.authRepo.GenerateToken(ctx, "secret", time.Hour*24, id)
	if err != nil {
		return "", errors.New("error authRepo: GenerateToken access")
	}

	return accessToken, nil
}
