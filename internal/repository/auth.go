package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/t-lunch/t-lunch-backend/internal/models"
	"github.com/t-lunch/t-lunch-backend/pkg/postgres"
)

type AuthRepository struct {
	db *postgres.DB
}

func NewAuthRepository(db *postgres.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(ctx context.Context, user *models.User) error {
	// TODO шифрование пароля

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	//     return err
	// }

	query := `
        INSERT INTO users (name, surname, tg, office, emoji, email, password)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id`
	row := r.db.Pool.QueryRow(ctx, query, user.Name, user.Surname, user.Tg, user.Office, user.Emoji, user.Email, user.Password)
	err := row.Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (int64, error) {
	var id int64
	query := `
        SELECT id
        FROM users
        WHERE email = $1`
	row := r.db.Pool.QueryRow(ctx, query, email)
	err := row.Scan(&id)
	if err == pgx.ErrNoRows {
		return 0, errors.New("Пользователь не найден")
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthRepository) CheckPassword(ctx context.Context, email, password string) (bool, error) {
	// TODO
	// var hashedPassword string
	var passwordDB string
	query := `SELECT password FROM users WHERE email = $1`
	row := r.db.Pool.QueryRow(ctx, query, email)
	err := row.Scan(&passwordDB)
	if err != nil {
		return false, err
	}
	// return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
	return passwordDB == password, nil
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

func (r *AuthRepository) ValidateRefreshToken(ctx context.Context, token string) (bool, error) {
	// TODO
	if token == "mock_refresh_token" { // Замените на реализацию JWT
		return true, nil
	}
	return false, errors.New("Неверный refresh-токен")
}
