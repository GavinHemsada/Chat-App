package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GavinHemsada/go-backend/internal/config"
	"github.com/GavinHemsada/go-backend/internal/database"
	"github.com/GavinHemsada/go-backend/internal/handlers"
	repository "github.com/GavinHemsada/go-backend/internal/repositories"
	"github.com/GavinHemsada/go-backend/internal/router"
	"github.com/GavinHemsada/go-backend/internal/services"
)

func main() {
	// Load config
	cfg := config.Load()
	fmt.Println(cfg.DBUser)

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	// Run migrations
	log.Println("Running database migrations...")
	database.RunMigrations(dsn)

	// Connect to database using sqlx
	log.Println("Connecting to database...")
	db, err := database.ConnectDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo, cfg.JWTSecret)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)

	// Setup router
	r := router.NewRouter(userHandler, cfg.JWTSecret)

	// Setup HTTP server
	port := cfg.ServerPort
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}