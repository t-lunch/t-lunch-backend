package service

import (
	"context"

	"github.com/t-lunch-backend/internal/repository"
	"github.com/t-lunch-backend/pkg/api/generated"
)

type TlunchService struct {
	userRepo  repository.UserRepository
	lunchRepo repository.LunchRepository
}

func NewTlunchService(userRepo repository.UserRepository, lunchRepo repository.LunchRepository) *TlunchService {
	return &TlunchService{userRepo: userRepo, lunchRepo: lunchRepo}
}

func (s *TlunchService) Registration(ctx context.Context, req *generated.RegistrationRequest) (*generated.User, error) {
	return s.userRepo.CreateUser(ctx, req)
}

func (s *TlunchService) GetProfile(ctx context.Context, req *generated.UserRequest) (*generated.User, error) {
	return s.userRepo.GetUser(ctx, req.UserId)
}

func (s *TlunchService) CreateLunch(ctx context.Context, req *generated.CreateLunchRequest) (*generated.Lunch, error) {
	return s.lunchRepo.CreateLunch(ctx, req)
}

func (s *TlunchService) GetLunch(ctx context.Context, req *generated.LunchRequest) (*generated.Lunch, error) {
	return s.lunchRepo.GetLunch(ctx, req.LunchId)
}

func (s *TlunchService) JoinLunch(ctx context.Context, req *generated.JoinLunchRequest) error {
	return s.lunchRepo.JoinLunch(ctx, req.UserId, req.LunchId)
}

func (s *TlunchService) LeaveLunch(ctx context.Context, req *generated.LeaveLunchRequest) error {
	return s.lunchRepo.LeaveLunch(ctx, req.UserId, req.LunchId)
}

func (s *TlunchService) GetUserLunchHistory(ctx context.Context, req *generated.UserRequest) ([]*generated.LunchFeedback, error) {
	return s.lunchRepo.GetUserLunchHistory(ctx, req.UserId)
}

func (s *TlunchService) RateLunch(ctx context.Context, req *generated.RateLunchRequest) error {
	return s.lunchRepo.RateLunch(ctx, req.UserId, req.LunchId, req.IsLiked)
}

func (s *TlunchService) GetLunchFeedback(ctx context.Context, req *generated.LunchRequest) ([]*generated.LunchFeedback, error) {
	return s.lunchRepo.GetLunchFeedback(ctx, req.LunchId)
}

func (s *TlunchService) GetLunches(ctx context.Context, req *generated.GetLunchesRequest) ([]*generated.Lunch, error) {
	return s.lunchRepo.GetLunches(ctx, req.UserId, req.Offset, req.Limit)
}

func (s *TlunchService) GetLunchesByDate(ctx context.Context, req *generated.GetLunchesByDateRequest) ([]*generated.Lunch, error) {
	return s.lunchRepo.GetLunchesByDate(ctx, req.UserId, req.Date, req.Offset, req.Limit)
}

func (s *TlunchService) GetLunchesByDateAndOffice(ctx context.Context, req *generated.GetLunchesByDateAndOfficeRequest) ([]*generated.Lunch, error) {
	return s.lunchRepo.GetLunchesByDateAndOffice(ctx, req.UserId, req.Date, req.Office, req.Offset, req.Limit)
}

func (s *TlunchService) GetLunchesByOffice(ctx context.Context, req *generated.GetLunchesByOfficeRequest) ([]*generated.Lunch, error) {
	return s.lunchRepo.GetLunchesByOffice(ctx, req.UserId, req.Office, req.Offset, req.Limit)
}
