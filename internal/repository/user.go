package repository

import (
	"context"

	"github.com/t-lunch/t-lunch-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) GetUserPasswordByEmail(ctx context.Context, email string) (int64, string, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Select("id", "password").
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		return 0, "", err
	}

	return user.ID, user.HashedPassword, nil
}

func (r *UserRepository) IsUserExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("email = ?", email).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
