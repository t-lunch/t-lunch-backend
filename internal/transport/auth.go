package transport

import (
	"context"

	"github.com/AlekSi/pointer"
	"github.com/t-lunch/t-lunch-backend/internal/models"
	tlunch "github.com/t-lunch/t-lunch-backend/pkg/api/generated"
)

type AuthService interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Login(ctx context.Context, email, password string) (string, string, error)
	Refresh(ctx context.Context, token string, userId int64) (string, error)
}

type AuthTransport struct {
	tlunch.UnimplementedTlunchServer
	authService AuthService
}

func NewAuthTransport(authService AuthService) *AuthTransport {
	return &AuthTransport{authService: authService}
}

func (t *AuthTransport) Register(ctx context.Context, request *tlunch.RegisterRequest) (*tlunch.User, error) {
	user := &models.User{
		Name:           request.GetName(),
		Surname:        request.GetSurname(),
		Tg:             request.GetTg(),
		Office:         request.GetOffice(),
		Emoji:          request.GetEmoji(),
		Email:          request.GetEmail(),
		HashedPassword: request.GetPassword(),
	}

	response, err := t.authService.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	rsafe := pointer.Get(response)
	return &tlunch.User{
		UserId:  rsafe.ID,
		Name:    rsafe.Name,
		Surname: rsafe.Surname,
		Tg:      rsafe.Tg,
		Office:  rsafe.Office,
		Emoji:   rsafe.Emoji,
	}, nil
}

func (t *AuthTransport) Login(ctx context.Context, request *tlunch.LoginRequest) (*tlunch.LoginResponse, error) {
	accessToken, refreshToken, err := t.authService.Login(ctx, request.GetEmail(), request.GetPassword())
	if err != nil {
		return nil, err
	}
	return &tlunch.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (t *AuthTransport) Refresh(ctx context.Context, request *tlunch.RefreshRequest) (*tlunch.RefreshResponse, error) {
	accessToken, err := t.authService.Refresh(ctx, request.GetRefreshToken(), request.GetUserId())
	if err != nil {
		return nil, err
	}
	return &tlunch.RefreshResponse{
		AccessToken: accessToken,
	}, nil
}
