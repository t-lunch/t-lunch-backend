package repository

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/t-lunch/t-lunch-backend/pkg/postgres"
)

type AuthRepository struct {
	db *postgres.DB
}

func NewAuthRepository(db *postgres.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GenerateToken(ctx context.Context, secret string, expiration time.Duration, id int64) (string, error) {
	payload := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(expiration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(secret))
}

func (r *AuthRepository) ValidateToken(ctx context.Context, secret, token string) (int64, bool) {
	jwtToken, err := jwt.ParseWithClaims(
		token,
		jwt.MapClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
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
