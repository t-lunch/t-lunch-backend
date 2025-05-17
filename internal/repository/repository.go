package repository

import (
	"errors"

	"github.com/t-lunch/t-lunch-backend/internal/config"
	"gorm.io/gorm"
)

type TLunchRepos struct {
	AuthRepo  *AuthRepository
	LunchRepo *LunchRepository
	UserRepo  *UserRepository
}

func NewTLunchRepos(cfg *config.Config, db *gorm.DB) (*TLunchRepos, error) {
	if cfg == nil {
		return nil, errors.New("cfg is nil")
	}
	if db == nil {
		return nil, errors.New("db is nil")
	}
	return &TLunchRepos{
		AuthRepo:  NewAuthRepository(cfg),
		LunchRepo: NewLunchRepository(db),
		UserRepo:  NewUserRepository(db),
	}, nil
}
