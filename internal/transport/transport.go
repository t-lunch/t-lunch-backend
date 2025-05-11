package transport

import (
	"context"
	"fmt"

	"github.com/t-lunch/t-lunch-backend/internal/service"
	tlunch "github.com/t-lunch/t-lunch-backend/pkg/api/generated"
)

type TLunchServer struct {
	tlunch.UnimplementedTlunchServer
	AuthTransport  *AuthTransport
	LunchTransport *LunchTransport
}

func NewTLunchServer(services *service.TLunchServices) *TLunchServer {
	return &TLunchServer{
		AuthTransport:  NewAuthTransport(services.AuthService),
		LunchTransport: NewLunchTransport(services.LunchService),
	}
}

func (t *TLunchServer) Registration(ctx context.Context, request *tlunch.RegistrationRequest) (*tlunch.User, error) {
	return t.AuthTransport.Registration(ctx, request)
}

func (t *TLunchServer) Login(ctx context.Context, request *tlunch.LoginRequest) (*tlunch.LoginResponse, error) {
	return t.AuthTransport.Login(ctx, request)
}

func (t *TLunchServer) Refresh(ctx context.Context, request *tlunch.RefreshRequest) (*tlunch.RefreshResponse, error) {
	return t.AuthTransport.Refresh(ctx, request)
}

func (t *TLunchServer) CreateLunch(ctx context.Context, request *tlunch.CreateLunchRequest) (*tlunch.LunchResponse, error) {
	fmt.Println("SERVER:", request)
	return t.LunchTransport.CreateLunch(ctx, request)
}

func (t *TLunchServer) GetLunches(ctx context.Context, request *tlunch.LunchRequest) (*tlunch.GetLunchesResponse, error) {
	return t.LunchTransport.GetLunches(ctx, request)
}
