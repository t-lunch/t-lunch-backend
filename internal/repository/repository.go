package repository

import (
	"github.com/t-lunch/t-lunch-backend/internal/config"
	"github.com/t-lunch/t-lunch-backend/internal/errors"
	"gorm.io/gorm"
)

type TLunchRepos struct {
	AuthRepo  *AuthRepository
	LunchRepo *LunchRepository
	UserRepo  *UserRepository
}

func NewTLunchRepos(cfg *config.Config, db *gorm.DB) (*TLunchRepos, error) {
	if cfg == nil {
		return nil, errors.ErrConfigIsNil
	}
	if db == nil {
		return nil, errors.ErrDBIsNil
	}
	return &TLunchRepos{
		AuthRepo:  NewAuthRepository(cfg),
		LunchRepo: NewLunchRepository(db),
		UserRepo:  NewUserRepository(db),
	}, nil
}
