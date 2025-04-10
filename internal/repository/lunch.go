package repository

import (
	"context"
	"t-lunch-backend/internal/models"

	"database/sql"
)

type PostgresLunchRepository struct {
	db *sql.DB
}

func NewLunchRepository(db *sql.DB) LunchRepository {
	return &PostgresLunchRepository{db: db}
}

func (r *PostgresLunchRepository) CreateLunch(ctx context.Context, lunch *models.Lunch) error {
	query := "INSERT INTO lunches (fields...) VALUES (?, ?, ...)"
	_, err := r.db.ExecContext(ctx, query, lunch.Field1, lunch.Field2, lunch.Field3)
	return err
}

func (r *PostgresLunchRepository) GetLunches(ctx context.Context, userID int64, offset, limit int) ([]models.Lunch, error) {
	var lunches []models.Lunch
	query := "SELECT * FROM lunches WHERE user_id = ? LIMIT ? OFFSET ?"
	rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var lunch models.Lunch
		if err := rows.Scan(&lunch.Field1, &lunch.Field2, &lunch.Field3); err != nil {
			return nil, err
		}
		lunches = append(lunches, lunch)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return lunches, err
}

func (r *PostgresLunchRepository) GetLunchByID(ctx context.Context, lunchID int64) (*models.Lunch, error) {
	var lunch models.Lunch
	query := "SELECT * FROM lunches WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, lunchID)
	err := row.Scan(&lunch.Field1, &lunch.Field2, &lunch.Field3)
	if err != nil {
		return nil, err
	}
	return &lunch, nil
}

func (r *PostgresLunchRepository) JoinLunch(ctx context.Context, lunchID, userID int64) error {
	query := "INSERT INTO lunch_participants (lunch_id, user_id) VALUES (?, ?)"
	_, err := r.db.ExecContext(ctx, query, lunchID, userID)
	return err
}

func (r *PostgresLunchRepository) LeaveLunch(ctx context.Context, lunchID, userID int64) error {
	query := "DELETE FROM lunch_participants WHERE lunch_id = ? AND user_id = ?"
	_, err := r.db.ExecContext(ctx, query, lunchID, userID)
	return err
}

func (r *PostgresLunchRepository) RateLunch(ctx context.Context, userID, lunchID int64, isLiked bool) error {
	query := "INSERT INTO lunch_feedback (user_id, lunch_id, is_liked) VALUES (?, ?, ?) ON CONFLICT (user_id, lunch_id) DO UPDATE SET is_liked = ?"
	_, err := r.db.ExecContext(ctx, query, userID, lunchID, isLiked, isLiked)
	return err
}

func (r *PostgresLunchRepository) GetLunchHistory(ctx context.Context, userID int64) ([]models.LunchFeedback, error) {
	var history []models.LunchFeedback
	query := "SELECT * FROM history WHERE user_id = ?"
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var feedback models.LunchFeedback
		if err := rows.Scan(&feedback.Field1, &feedback.Field2, &feedback.Field3); err != nil {
			return nil, err
		}
		history = append(history, feedback)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return history, nil
}
