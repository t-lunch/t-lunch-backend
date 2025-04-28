package repository

import (
	"context"

	"github.com/t-lunch/t-lunch-backend/internal/models"
	"github.com/t-lunch/t-lunch-backend/pkg/postgres"
)

type UserRepository struct {
	db *postgres.DB
}

func NewUserRepository(db *postgres.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `
        INSERT INTO users (name, surname, tg, office, emoji, email, password)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id`
	row := r.db.Pool.QueryRow(ctx, query, user.Name, user.Surname, user.Tg, user.Office, user.Emoji, user.Email, user.HashedPassword)
	err := row.Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserPasswordByEmail(ctx context.Context, email string) (int64, string, error) {
	var id int64
	var hashedPassword string
	query := `
        SELECT id, password
        FROM users
        WHERE email = $1`
	row := r.db.Pool.QueryRow(ctx, query, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil
}
