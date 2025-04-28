package repository

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/t-lunch/t-lunch-backend/internal/config"
	"github.com/t-lunch/t-lunch-backend/internal/models"
)

type AuthRepository struct {
	secret            string
	accessExpiration  time.Duration
	refreshExpiration time.Duration
}

func NewAuthRepository(cfg *config.Config) *AuthRepository {
	return &AuthRepository{
		secret:            cfg.Jwt.Secret,
		accessExpiration:  time.Hour * time.Duration(cfg.Jwt.AccessExpiration),
		refreshExpiration: time.Hour * 24 * time.Duration(cfg.Jwt.RefreshExpiration),
	}
}

func (r *AuthRepository) GenerateToken(ctx context.Context, id int64, tokenType models.TokenType) (string, error) {
	payload := jwt.MapClaims{
		"id": id,
	}

	switch tokenType {
	case models.Access:
		payload["exp"] = time.Now().Add(r.accessExpiration).Unix()
	case models.Refresh:
		payload["exp"] = time.Now().Add(r.refreshExpiration).Unix()
	default:
		return "", errors.New("unknown token type")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(r.secret))
}

func (r *AuthRepository) ValidateToken(ctx context.Context, token string) (int64, bool) {
	jwtToken, err := jwt.ParseWithClaims(
		token,
		jwt.MapClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(r.secret), nil
		},
	)
	if err != nil {
		return -1, false
	}

	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return -2, ok
	}

	if int64(payload["exp"].(float64)) < time.Now().Unix() {
		return -3, false
	}

	return int64(payload["id"].(float64)), true
}
