package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/t-lunch/t-lunch-backend/internal/errors"
	"github.com/t-lunch/t-lunch-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserPasswordByEmail(ctx context.Context, email string) (int64, string, error)
	IsUserWithEmailExist(ctx context.Context, email string) (bool, error)
}

type AuthRepo interface {
	GenerateToken(ctx context.Context, id int64, tokenType models.TokenType) (string, error)
	GetToken(ctx context.Context, token string) (jwt.MapClaims, error)
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

func (s *AuthService) Register(ctx context.Context, user *models.User) (*models.User, error) {
	if user.Name == "" || user.Surname == "" || user.Tg == "" || user.Office == "" || user.Emoji == "" || user.Email == "" || user.HashedPassword == "" {
		return nil, errors.ErrInvalidRequest
	}

	exists, err := s.userRepo.IsUserWithEmailExist(ctx, user.Email)
	if err != nil {
		return nil, errors.NewErrRepository("userRepo", "IsUserWithEmailExist", err)
	}
	if exists {
		return nil, errors.NewErrUserWithEmailAlreadyExists(user.Email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.HashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.NewErrRepository("bcrypt", "GenerateFromPassword", err)
	}
	user.HashedPassword = string(hashedPassword)

	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.NewErrRepository("userRepo", "CreateUser", err)
	}

	return user, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, string, error) {
	if email == "" || password == "" {
		return "", "", errors.ErrInvalidRequest
	}

	id, hashedPassword, err := s.userRepo.GetUserPasswordByEmail(ctx, email)
	if err != nil {
		return "", "", errors.NewErrRepository("userRepo", "GetUserPasswordByEmail", err)
	}

	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return "", "", errors.NewErrRepository("bcrypt", "CompareHashAndPassword", errors.ErrInvalidPassword)
	}

	accessToken, err := s.authRepo.GenerateToken(ctx, id, models.Access)
	if err != nil {
		return "", "", errors.NewErrRepository("authRepo", "GenerateToken access", err)
	}

	refreshToken, err := s.authRepo.GenerateToken(ctx, id, models.Refresh)
	if err != nil {
		return "", "", errors.NewErrRepository("authRepo", "GenerateToken refresh", err)
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) Refresh(ctx context.Context, token string, userId int64) (string, error) {
	if token == "" {
		return "", errors.ErrInvalidRequest
	}

	payload, err := s.authRepo.GetToken(ctx, token)
	if err != nil {
		return "", errors.NewErrRepository("authRepo", "GetToken", err)
	}

	if int64(payload["exp"].(float64)) < time.Now().Unix() {
		return "", errors.NewErrRepository("authRepo", "ValidateToken", errors.ErrTokenExpired)
	}

	id := int64(payload["id"].(float64))
	if id != userId {
		return "", errors.NewErrRepository("authRepo", "GetToken", errors.NewErrUserAndOwnerAreDifferent(userId, id))
	}

	accessToken, err := s.authRepo.GenerateToken(ctx, id, models.Access)
	if err != nil {
		return "", errors.NewErrRepository("authRepo", "GenerateToken access", err)
	}

	return accessToken, nil
}
