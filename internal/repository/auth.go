package repository

import (
	"context"

	"github.com/t-lunch/t-lunch-backend/pkg/postgres"
)

type AuthRepository struct {
	db *postgres.DB
}

func NewAuthRepository(db *postgres.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GenerateAccessToken(ctx context.Context) string {
	// TODO
	accessToken := "mock_access_token" // Замените на реализацию JWT
	return accessToken
}

func (r *AuthRepository) GenerateRefreshToken(ctx context.Context) string {
	// TODO
	refreshToken := "mock_refresh_token" // Замените на реализацию JWT
	return refreshToken
}
