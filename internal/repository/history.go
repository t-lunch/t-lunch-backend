package repository

import (
	"context"

	"github.com/t-lunch-backend/internal/models"

	"database/sql"
)

type PostgresHistoryRepository struct {
	db *sql.DB
}

func NewHistoryRepository(db *sql.DB) HistoryRepository {
	return &PostgresHistoryRepository{db: db}
}

func (r *PostgresHistoryRepository) CreateHistory(ctx context.Context, history *models.History) error {
	query := "INSERT INTO histories (user_id, lunch_id, is_liked) VALUES ($1, $2, $3) RETURNING id"
	return r.db.QueryRowContext(ctx, query, history.UserID, history.LunchID, history.IsLiked).Scan(&history.ID)
}

func (r *PostgresHistoryRepository) GetHistoryByUser(ctx context.Context, userID int64) ([]models.History, error) {
	query := "SELECT id, user_id, lunch_id, is_liked FROM histories WHERE user_id = $1"
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []models.History
	for rows.Next() {
		var history models.History
		if err := rows.Scan(&history.ID, &history.UserID, &history.LunchID, &history.IsLiked); err != nil {
			return nil, err
		}
		histories = append(histories, history)
	}
	return histories, nil
}

func (r *PostgresHistoryRepository) UpdateHistory(ctx context.Context, history *models.History) error {
	query := "UPDATE histories SET is_liked = $1 WHERE id = $2"
	_, err := r.db.ExecContext(ctx, query, history.IsLiked, history.ID)
	return err
}
