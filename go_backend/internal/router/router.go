package router

import (
	"net/http"

	"github.com/GavinHemsada/go-backend/internal/handlers"
	"github.com/GavinHemsada/go-backend/internal/middleware"
	"github.com/GavinHemsada/go-backend/internal/websocket"
	"github.com/gorilla/mux"
)

func NewRouter(userHandler *handlers.UserHandler, roomHandler *handlers.RoomHandler, messageHandler *handlers.MessageHandler, wsHandler *websocket.Handler, jwtSecret string) *mux.Router {
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
	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.JWTMiddleware(jwtSecret))
	
	// User routes
	users := protected.PathPrefix("/users").Subrouter()
	users.HandleFunc("/{id}", userHandler.GetUserByID).Methods("GET")
	users.HandleFunc("", userHandler.GetAllUsers).Methods("GET")
	
	// Room routes
	rooms := protected.PathPrefix("/rooms").Subrouter()
	rooms.HandleFunc("", roomHandler.CreateRoom).Methods("POST")
	rooms.HandleFunc("", roomHandler.GetAllRooms).Methods("GET")
	rooms.HandleFunc("/user", roomHandler.GetUserRooms).Methods("GET")
	rooms.HandleFunc("/{id}", roomHandler.GetRoomByID).Methods("GET")
	rooms.HandleFunc("/{id}", roomHandler.DeleteRoom).Methods("DELETE")
	rooms.HandleFunc("/{id}/join", roomHandler.JoinRoom).Methods("POST")
	rooms.HandleFunc("/{id}/leave", roomHandler.LeaveRoom).Methods("POST")
	rooms.HandleFunc("/{id}/members", roomHandler.GetRoomMembers).Methods("GET")
	
	// Message routes
	messages := protected.PathPrefix("/rooms/{room_id}/messages").Subrouter()
	messages.HandleFunc("", messageHandler.CreateMessage).Methods("POST")
	messages.HandleFunc("", messageHandler.GetMessagesByRoom).Methods("GET")
	
	// WebSocket route for live chat (protected with JWT)
	protected.HandleFunc("/ws/rooms/{room_id}", wsHandler.ServeWS)

	return r
}
