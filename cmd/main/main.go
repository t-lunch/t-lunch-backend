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
	"github.com/t-lunch/t-lunch-backend/internal/transport"
	"github.com/t-lunch/t-lunch-backend/pkg/gorm"
)

var cfgName string = "lunch"

func main() {
	cfg, err := config.NewConfig(cfgName)
	if err != nil {
		fmt.Println("error: NewCfg")
		return
	}

	gormDB, err := gorm.NewGormDB(
		context.Background(),
		fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port),
	)
	if err != nil {
		fmt.Println("error: NewDB")
		return
	}

	repos, err := repository.NewTLunchRepos(cfg, gormDB)
	if err != nil {
		fmt.Println("error: repos")
		return
	}
	services := service.NewTLunchServices(repos)
	transports := transport.NewTLunchServer(services)
	server := app.NewServer(cfg.Server.HTTPPort, cfg.ProtectedUrl, transports)

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
