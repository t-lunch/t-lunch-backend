package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/t-lunch/t-lunch-backend/internal/transport"
	tlunch "github.com/t-lunch/t-lunch-backend/pkg/api/generated"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	restServer *http.Server
}

func NewServer(restPort int, service transport.AuthService) *Server {
	authSrv := transport.NewAuthTransport(service)

	restServer := runtime.NewServeMux()
	if err := tlunch.RegisterTlunchHandlerServer(context.Background(), restServer, authSrv); err != nil {
		fmt.Printf("Failed to listen rest: %s\n", err)
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", restPort),
		Handler: restServer,
	}

	return &Server{restServer: httpServer}
}

func (s *Server) Start() error {
	eg := errgroup.Group{}

	eg.Go(func() error {
		return s.restServer.ListenAndServe()
	})

	return eg.Wait()
}

func (s *Server) Stop() error {
	return s.restServer.Shutdown(context.Background())
}
