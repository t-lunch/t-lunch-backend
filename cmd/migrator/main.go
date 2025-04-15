package main

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/t-lunch/t-lunch-backend/internal/config"
)

var cfgName string = "lunch"

func main() {
	cfg, err := config.NewConfig(cfgName)
	if err != nil {
		fmt.Println("error: NewCfg")
		return
	}

	dsn := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", cfg.Database.User, cfg.Database.Password, cfg.Database.Port, cfg.Database.DBName)

	fmt.Println(dsn)

	m, err := migrate.New("file://internal/migrations", dsn)
	if err != nil {
		fmt.Printf("panic error1: %v\n", err)
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migration to change")
			return
		}
		fmt.Printf("panic error2: %v\n", err)
		panic(err)
	}

	fmt.Println("migrations applied successfully")
}
