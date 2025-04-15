package repository

// import (
// 	"context"

// 	"github.com/t-lunch/t-lunch-backend/internal/models"

// 	"database/sql"
// )

// type historyRepository struct {
// 	db *sql.DB
// }

// func NewHistoryRepository(db *sql.DB) HistoryRepository {
// 	return &historyRepository{db: db}
// }

// func (r *historyRepository) CreateHistory(ctx context.Context, history *models.History) error {
// 	query := "INSERT INTO histories (user_id, lunch_id, is_liked) VALUES ($1, $2, $3) RETURNING id"
// 	return r.db.QueryRowContext(ctx, query, history.UserID, history.LunchID, history.IsLiked).Scan(&history.ID)
// }

// func (r *historyRepository) RateLunch(ctx context.Context, userID, lunchID int64, isLiked bool) error {
// 	query := `INSERT INTO history (user_id, lunch_id, is_liked)
// 	          VALUES ($1, $2, $3)
// 	          ON CONFLICT (user_id, lunch_id)
// 	          DO UPDATE SET is_liked = EXCLUDED.is_liked`
// 	_, err := r.db.ExecContext(ctx, query, userID, lunchID, isLiked)
// 	return err
// }

// func (r *historyRepository) GetHistoryByUser(ctx context.Context, userID int64) ([]models.History, error) {
// 	query := "SELECT id, user_id, lunch_id, is_liked FROM histories WHERE user_id = $1"
// 	rows, err := r.db.QueryContext(ctx, query, userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var histories []models.History
// 	for rows.Next() {
// 		var history models.History
// 		if err := rows.Scan(&history.ID, &history.UserID, &history.LunchID, &history.IsLiked); err != nil {
// 			return nil, err
// 		}
// 		histories = append(histories, history)
// 	}
// 	return histories, nil
// }
