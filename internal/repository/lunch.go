package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/t-lunch/t-lunch-backend/internal/models"
	"gorm.io/gorm"
)

type LunchRepository struct {
	db *gorm.DB
}

func NewLunchRepository(db *gorm.DB) *LunchRepository {
	return &LunchRepository{db: db}
}

func (r *LunchRepository) CreateLunch(ctx context.Context, lunch *models.Lunch) error {
	return r.db.WithContext(ctx).Preload("Creator").Create(lunch).Error
}

func (r *LunchRepository) GetLunches(ctx context.Context, userID int64, offset, limit int) ([]*models.Lunch, error) {
	var lunches []*models.Lunch
	err := r.db.WithContext(ctx).
		Preload("Creator").
		Order("time DESC").
		Offset(offset).
		Limit(limit).
		Find(&lunches).Error
	if err != nil {
		return nil, err
	}

	return lunches, nil
}

func (r *LunchRepository) GetLunchByID(ctx context.Context, id int64) (*models.Lunch, error) {
	var lunch models.Lunch
	err := r.db.WithContext(ctx).Preload("Creator").First(&lunch, id).Error
	if err != nil {
		return nil, err
	}

	return &lunch, nil
}

func (r *LunchRepository) GetLunchIdByUserID(ctx context.Context, userID int64) (int64, error) {
	var lunch models.Lunch
	err := r.db.WithContext(ctx).
		Select("id").
		Where("? = ANY(participants)", userID).
		First(&lunch).Error
	if err != nil {
		return 0, err
	}

	return lunch.ID, nil
}

func (r *LunchRepository) UpdateLunchParticipants(ctx context.Context, method models.UpdateAction, lunchID, userID int64) error {
	expr := fmt.Sprintf("%s(participants, ?)", method)
	err := r.db.WithContext(ctx).
		Model(&models.Lunch{}).
		Where("id = ?", lunchID).
		Update("participants", gorm.Expr(expr, userID)).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *LunchRepository) GetUsersLunches(ctx context.Context, userID int64, offset, limit int) ([]*models.Lunch, error) {
	var lunches []*models.Lunch
	err := r.db.WithContext(ctx).
		Where("? = ANY(participants)", userID).
		Where("time < ?", time.Now()).
		Preload("Creator").
		Order("time DESC").
		Offset(offset).
		Limit(limit).
		Find(&lunches).Error
	if err != nil {
		return nil, err
	}

	return lunches, nil
}

func (r *LunchRepository) UpdateLunchLikedBy(ctx context.Context, method models.UpdateAction, lunchID, userID int64) error {
	expr := fmt.Sprintf("%s(liked_by, ?)", method)
	err := r.db.WithContext(ctx).
		Model(&models.Lunch{}).
		Where("id = ?", lunchID).
		Update("liked_by", gorm.Expr(expr, userID)).Error
	if err != nil {
		return err
	}

	return nil
}
