package repository

import (
	"context"

	"github.com/t-lunch-backend/internal/models"

	"database/sql"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password)
	return err
}

func (r *PostgresUserRepository) GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	var user models.User
	query := "SELECT user_id, name, email, password FROM users WHERE user_id = $1"
	row := r.db.QueryRowContext(ctx, query, userID)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := "SELECT user_id, name, email, password FROM users WHERE email = $1"
	row := r.db.QueryRowContext(ctx, query, email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	query := "UPDATE users SET name = $1, email = $2, password = $3 WHERE user_id = $4"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.ID)
	return err
}
