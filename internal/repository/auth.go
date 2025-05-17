package repository

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/t-lunch/t-lunch-backend/internal/config"
	"github.com/t-lunch/t-lunch-backend/internal/errors"
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
		return "", errors.ErrUnknownTokenType
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(r.secret))
}

func (r *AuthRepository) GetToken(ctx context.Context, token string) (jwt.MapClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(
		token,
		jwt.MapClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(r.secret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.ErrInvalidTokenClaims
	}

	return payload, nil
}
