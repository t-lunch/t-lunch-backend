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
	"github.com/t-lunch/t-lunch-backend/pkg/logger"
	"go.uber.org/zap"
)

var cfgName string = "lunch"

func main() {
	zapLogger, err := logger.NewZapLogger()
	if err != nil {
		fmt.Printf("failed to create new logger: %v\n", err)
		return
	}

	defer zapLogger.Sync()

	cfg, err := config.NewConfig(cfgName)
	if err != nil {
		zapLogger.Error("failed to initialize config", zap.Error(err))
		return
	}

	gormDB, err := gorm.NewGormDB(
		context.Background(),
		fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port),
	)
	if err != nil {
		zapLogger.Error("failed to connect to DB", zap.Error(err))
		return
	}

	repos, err := repository.NewTLunchRepos(cfg, gormDB)
	if err != nil {
		zapLogger.Error("failed to initialize TLunch repositories", zap.Error(err))
		return
	}
	services := service.NewTLunchServices(repos)
	transports := transport.NewTLunchServer(services, zapLogger)
	server := app.NewServer(cfg.Server.HTTPPort, cfg.ProtectedUrl, cfg.Jwt.Secret, transports)

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGTERM, syscall.SIGINT)

	zapLogger.Info("server started", zap.String("server", "TLunch"), zap.Int("port", cfg.Server.HTTPPort))

	go func() {
		if err := server.Start(); err != nil {
			zapLogger.Error("failed to start", zap.Error(err))
		}
	}()

	<-gracefulShutdown

	if err := server.Stop(); err != nil {
		zapLogger.Error("failed to graceful stop", zap.Error(err))
	}

	zapLogger.Info("server stopped graceful", zap.String("server", "TLunch"), zap.Int("port", cfg.Server.HTTPPort))
}
