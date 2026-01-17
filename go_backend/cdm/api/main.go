package main

import (
	"context"
	"fmt"
	"log"

	"github.com/GavinHemsada/go-backend/internal/config"
	"github.com/GavinHemsada/go-backend/internal/database"
)

func main() {
	// Load config
	cfg := config.Load()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	// AUTO MIGRATION
	database.RunMigrations(dsn)

	// Connect to DB
	db := database.Connect(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	defer db.Close(context.Background())

	log.Println("App is ready!")
}