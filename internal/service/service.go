package service

import "github.com/t-lunch/t-lunch-backend/internal/repository"

type TLunchServices struct {
	AuthService  *AuthService
	LunchService *LunchService
}

func NewTLunchServices(repos *repository.TLunchRepos) *TLunchServices {
	return &TLunchServices{
		AuthService:  NewAuthService(repos.UserRepo, repos.AuthRepo),
		LunchService: NewLunchService(repos.LunchRepo),
	}
}
