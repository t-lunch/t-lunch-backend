package transport

import (
	"context"

	"github.com/AlekSi/pointer"
	"github.com/t-lunch/t-lunch-backend/internal/models"
	tlunch "github.com/t-lunch/t-lunch-backend/pkg/api/generated"
	"go.uber.org/zap"
)

type AuthService interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Login(ctx context.Context, email, password string) (string, string, error)
	Refresh(ctx context.Context, token string, userId int64) (string, error)
}

type AuthTransport struct {
	tlunch.UnimplementedTlunchServer
	authService AuthService
	zapLogger   *zap.Logger
}

func NewAuthTransport(authService AuthService, zapLogger *zap.Logger) *AuthTransport {
	return &AuthTransport{
		authService: authService,
		zapLogger:   zapLogger,
	}
}

func (t *AuthTransport) Register(ctx context.Context, request *tlunch.RegisterRequest) (*tlunch.User, error) {
	t.zapLogger.Info("Register request", zap.String("email", request.GetEmail()))

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
		t.zapLogger.Error("Register failed", zap.Error(err))
		return nil, err
	}

	rsafe := pointer.Get(response)
	t.zapLogger.Info("Register success", zap.String("email", request.GetEmail()), zap.Int64("user_id", rsafe.ID))

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
	t.zapLogger.Info("Login request", zap.String("email", request.GetEmail()))

	accessToken, refreshToken, err := t.authService.Login(ctx, request.GetEmail(), request.GetPassword())
	if err != nil {
		t.zapLogger.Error("Login failed", zap.Error(err))
		return nil, err
	}

	t.zapLogger.Info("Login success", zap.String("email", request.GetEmail()))

	return &tlunch.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (t *AuthTransport) Refresh(ctx context.Context, request *tlunch.RefreshRequest) (*tlunch.RefreshResponse, error) {
	t.zapLogger.Info("Refresh request", zap.Int64("user_id", request.GetUserId()))

	accessToken, err := t.authService.Refresh(ctx, request.GetRefreshToken(), request.GetUserId())
	if err != nil {
		t.zapLogger.Error("Refresh failed", zap.Error(err))
		return nil, err
	}

	t.zapLogger.Info("Refresh success", zap.Int64("user_id", request.GetUserId()))

	return &tlunch.RefreshResponse{
		AccessToken: accessToken,
	}, nil
}
