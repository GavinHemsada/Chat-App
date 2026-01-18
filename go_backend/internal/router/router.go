package router

import (
	"net/http"

	"github.com/GavinHemsada/go-backend/internal/handlers"
	"github.com/GavinHemsada/go-backend/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(userHandler *handlers.UserHandler, jwtSecret string) *mux.Router {
	r := mux.NewRouter()

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()
	
	// Public routes (no authentication required)
	api.HandleFunc("/users/register", userHandler.Register).Methods("POST")
	api.HandleFunc("/users/login", userHandler.Login).Methods("POST")
	
	// Protected routes (require JWT authentication)
	protected := api.PathPrefix("/users").Subrouter()
	protected.Use(middleware.JWTMiddleware(jwtSecret))
	// These routes require authentication
	protected.HandleFunc("/{id}", userHandler.GetUserByID).Methods("GET")
	protected.HandleFunc("", userHandler.GetAllUsers).Methods("GET")

	return r
}
