package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/t-lunch/t-lunch-backend/internal/app"
	"github.com/t-lunch/t-lunch-backend/internal/config"
	"github.com/t-lunch/t-lunch-backend/internal/repository"
	"github.com/t-lunch/t-lunch-backend/internal/service"
	"github.com/t-lunch/t-lunch-backend/pkg/postgres"
)

var cfgName string = "lunch"

func main() {
	cfg, err := config.NewConfig(cfgName)
	if err != nil {
		fmt.Println("error: NewCfg")
		return
	}

	psql, err := postgres.NewDB(
		context.Background(),
		fmt.Sprintf("postgres://%s:%s@postgres:%s/%s?sslmode=disable", cfg.Database.User, cfg.Database.Password, cfg.Database.Port, cfg.Database.DBName),
	)
	if err != nil {
		fmt.Println("error: NewDB")
		return
	}

	// fmt.Println("KAIF")

	userRepo := repository.NewUserRepository(psql)
	authRepo := repository.NewAuthRepository(cfg)
	srv := service.NewAuthService(userRepo, authRepo)
	server := app.NewServer(cfg.Server.HTTPPort, srv)

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := server.Start(); err != nil {
			fmt.Printf("Failed to start: %s\n", err)
		}
	}()

	<-gracefulShutdown

	if err := server.Stop(); err != nil {
		fmt.Printf("Failed to graceful stop: %s\n", err)
	}

	fmt.Println("Server Stopped")
}
