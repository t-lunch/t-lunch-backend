package transport

import (
	"context"

	"github.com/AlekSi/pointer"
	"github.com/t-lunch/t-lunch-backend/internal/models"
	tlunch "github.com/t-lunch/t-lunch-backend/pkg/api/generated"
	"go.uber.org/zap"
)

type UserService interface {
	GetUsersByIDs(ctx context.Context, userIDs []int64) ([]*models.UserResponse, error)
	GetUserByID(ctx context.Context, userID int64) (*models.UserResponse, error)
	UpdateUser(ctx context.Context, user *models.UserResponse) (*models.UserResponse, error)
}

type UserTransport struct {
	tlunch.UnimplementedTlunchServer
	userService UserService
	zapLogger   *zap.Logger
}

func NewUserTransport(userService UserService, zapLogger *zap.Logger) *UserTransport {
	return &UserTransport{
		userService: userService,
		zapLogger:   zapLogger,
	}
}

func (t *UserTransport) GetProfile(ctx context.Context, request *tlunch.UserRequest) (*tlunch.User, error) {
	t.zapLogger.Info("GetProfile request", zap.Int64("user_id", request.GetUserId()))

	user, err := t.userService.GetUserByID(ctx, request.GetUserId())
	if err != nil {
		t.zapLogger.Error("GetProfile failed", zap.Error(err))
		return nil, err
	}

	rsafe := pointer.Get(user)
	t.zapLogger.Info("GetProfile success", zap.Int64("user_id", user.ID))

	return &tlunch.User{
		UserId:  rsafe.ID,
		Name:    rsafe.Name,
		Surname: rsafe.Surname,
		Tg:      rsafe.Tg,
		Office:  rsafe.Office,
		Emoji:   rsafe.Emoji,
	}, nil
}

func (t *UserTransport) ChangeProfile(ctx context.Context, request *tlunch.User) (*tlunch.User, error) {
	t.zapLogger.Info("ChangeProfile request", zap.Int64("user_id", request.GetUserId()))

	user, err := t.userService.UpdateUser(ctx, &models.UserResponse{
		ID:      request.GetUserId(),
		Name:    request.GetName(),
		Surname: request.GetSurname(),
		Tg:      request.GetTg(),
		Office:  request.GetOffice(),
		Emoji:   request.GetEmoji(),
	})
	if err != nil {
		t.zapLogger.Error("ChangeProfile failed", zap.Error(err))
		return nil, err
	}

	rsafe := pointer.Get(user)
	t.zapLogger.Info("ChangeProfile success", zap.Int64("user_id", user.ID))

	return &tlunch.User{
		UserId:  rsafe.ID,
		Name:    rsafe.Name,
		Surname: rsafe.Surname,
		Tg:      rsafe.Tg,
		Office:  rsafe.Office,
		Emoji:   rsafe.Emoji,
	}, nil
}
