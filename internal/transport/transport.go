package transport

import (
	"context"

	"github.com/t-lunch/t-lunch-backend/internal/service"
	tlunch "github.com/t-lunch/t-lunch-backend/pkg/api/generated"
	"go.uber.org/zap"
)

type TLunchServer struct {
	tlunch.UnimplementedTlunchServer
	AuthTransport  *AuthTransport
	LunchTransport *LunchTransport
	UserTransport  *UserTransport
}

func NewTLunchServer(services *service.TLunchServices, zapLogger *zap.Logger) *TLunchServer {
	return &TLunchServer{
		AuthTransport:  NewAuthTransport(services.AuthService, zapLogger),
		LunchTransport: NewLunchTransport(services.LunchService, services.UserService, zapLogger),
		UserTransport:  NewUserTransport(services.UserService, zapLogger),
	}
}

func (t *TLunchServer) Register(ctx context.Context, request *tlunch.RegisterRequest) (*tlunch.User, error) {
	return t.AuthTransport.Register(ctx, request)
}

func (t *TLunchServer) Login(ctx context.Context, request *tlunch.LoginRequest) (*tlunch.LoginResponse, error) {
	return t.AuthTransport.Login(ctx, request)
}

func (t *TLunchServer) Refresh(ctx context.Context, request *tlunch.RefreshRequest) (*tlunch.RefreshResponse, error) {
	return t.AuthTransport.Refresh(ctx, request)
}

func (t *TLunchServer) CreateLunch(ctx context.Context, request *tlunch.CreateLunchRequest) (*tlunch.LunchResponse, error) {
	return t.LunchTransport.CreateLunch(ctx, request)
}

func (t *TLunchServer) GetLunches(ctx context.Context, request *tlunch.LunchRequest) (*tlunch.GetLunchesResponse, error) {
	return t.LunchTransport.GetLunches(ctx, request)
}

func (t *TLunchServer) GetDetailLunch(ctx context.Context, request *tlunch.DetailLunchRequest) (*tlunch.DetailLunchResponse, error) {
	return t.LunchTransport.GetDetailLunch(ctx, request)
}

func (t *TLunchServer) JoinLunch(ctx context.Context, request *tlunch.ActionLunchRequest) (*tlunch.LunchResponse, error) {
	return t.LunchTransport.JoinLunch(ctx, request)
}

func (t *TLunchServer) LeaveLunch(ctx context.Context, request *tlunch.ActionLunchRequest) (*tlunch.LunchResponse, error) {
	return t.LunchTransport.LeaveLunch(ctx, request)
}

func (t *TLunchServer) GetProfile(ctx context.Context, request *tlunch.UserRequest) (*tlunch.User, error) {
	return t.UserTransport.GetProfile(ctx, request)
}

func (t *TLunchServer) ChangeProfile(ctx context.Context, request *tlunch.User) (*tlunch.User, error) {
	return t.UserTransport.ChangeProfile(ctx, request)
}
