package repository

import (
	"context"

	"github.com/t-lunch-backend/internal/models"

	"database/sql"
)

type lunchRepository struct {
	db *sql.DB
}

func NewLunchRepository(db *sql.DB) LunchRepository {
	return &lunchRepository{db: db}
}

func (r *lunchRepository) CreateLunch(ctx context.Context, lunch *models.Lunch) error {
	query := `INSERT INTO lunches (name, surname, place, time, description, number_of_participants) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, lunch.Name, lunch.Surname, lunch.Place, lunch.Time, lunch.Description, lunch.NumberOfParticipants).Scan(&lunch.ID)
	return err
}

func (r *lunchRepository) GetLunches(ctx context.Context, userID int64, offset, limit int) ([]models.Lunch, error) {
	var lunches []models.Lunch
	query := `SELECT id, name, surname, place, time, number_of_participants, description 
	          FROM lunches ORDER BY time DESC LIMIT $1 OFFSET $2`
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

func (r *lunchRepository) GetLunchByID(ctx context.Context, lunchID int64) (*models.Lunch, error) {
	var lunch models.Lunch
	query := "SELECT * FROM lunches WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, lunchID)
	err := row.Scan(&lunch.Field1, &lunch.Field2, &lunch.Field3)
	if err != nil {
		return nil, err
	}
	return &lunch, nil
}

func (r *lunchRepository) JoinLunch(ctx context.Context, lunchID, userID int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Проверяем, не добавлен ли уже пользователь
	checkQuery := `SELECT COUNT(*) FROM lunch_participants WHERE lunch_id = $1 AND user_id = $2`
	var count int
	err = tx.QueryRowContext(ctx, checkQuery, lunchID, userID).Scan(&count)
	if err != nil {
		tx.Rollback()
		return err
	}

	if count == 0 {
		insertQuery := `INSERT INTO lunch_participants (lunch_id, user_id) VALUES ($1, $2)`
		_, err = tx.ExecContext(ctx, insertQuery, lunchID, userID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *lunchRepository) LeaveLunch(ctx context.Context, lunchID, userID int64) error {
	query := "DELETE FROM lunches WHERE lunch_id = ? AND user_id = ?"
	_, err := r.db.ExecContext(ctx, query, lunchID, userID)
	return err
}

/*func (r *lunchRepository) RateLunch(ctx context.Context, userID, lunchID int64, isLiked bool) error {
	query := "INSERT INTO histories (user_id, lunch_id, is_liked) VALUES (?, ?, ?) ON CONFLICT (user_id, lunch_id) DO UPDATE SET is_liked = ?"
	_, err := r.db.ExecContext(ctx, query, userID, lunchID, isLiked, isLiked)
	return err
}*/

func (r *lunchRepository) GetLunchHistory(ctx context.Context, userID int64) ([]models.LunchFeedback, error) {
	var history []models.LunchFeedback
	query := "SELECT * FROM histories WHERE user_id = ?"
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
