package websocket

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/GavinHemsada/go-backend/internal/middleware"
	"github.com/GavinHemsada/go-backend/internal/models"
	"github.com/GavinHemsada/go-backend/internal/services"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Configure properly for production
	},
}

type Handler struct {
	hub            *Hub
	messageService *services.MessageService
	roomService    *services.RoomService
}

func (h *Handler) GetHub() *Hub {
	return h.hub
}

func NewHandler(messageService *services.MessageService, roomService *services.RoomService, redisClient *redis.Client) *Handler {
	// Create message processor
	processor := func(msg *BroadcastMessage) *BroadcastMessage {
		ctx := context.Background()
		
		var wsMsg models.WSMessage
		if err := json.Unmarshal(msg.Message, &wsMsg); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			return nil
		}

		if wsMsg.Type != "message" {
			return msg // Forward non-message types as-is
		}

		// Parse user and room IDs
		userID, err := uuid.Parse(msg.UserID)
		if err != nil {
			log.Printf("Invalid user ID: %v", err)
			return nil
		}

		roomID, err := uuid.Parse(wsMsg.RoomID)
		if err != nil {
			log.Printf("Invalid room ID: %v", err)
			return nil
		}

		content := wsMsg.Content
		if content == "" {
			if payloadStr, ok := wsMsg.Payload.(string); ok {
				content = payloadStr
			}
		}

		if content == "" {
			log.Printf("Empty message content")
			return nil
		}

		// Save message to database
		savedMsg, err := messageService.CreateMessage(ctx, roomID, userID, content, "text")
		if err != nil {
			log.Printf("Error saving message: %v", err)
			return nil
		}

		// Create response with saved message
		response := models.WSMessageResponse{
			Type:    "message",
			Message: savedMsg,
			UserID:  userID.String(),
			RoomID:  roomID.String(),
		}

		responseBytes, err := json.Marshal(response)
		if err != nil {
			log.Printf("Error marshaling response: %v", err)
			return nil
		}

		// Return processed message (without UserID so it won't be processed again)
		return &BroadcastMessage{
			RoomID:    msg.RoomID,
			Message:   responseBytes,
			UserID:    "",      // Clear UserID to indicate it's processed
			FromRedis: false,   // This is a new message, not from Redis
		}
	}

	hub := NewHub(redisClient, processor)

	return &Handler{
		hub:            hub,
		messageService: messageService,
		roomService:    roomService,
	}
}

func (h *Handler) ServeWS(w http.ResponseWriter, r *http.Request) {
	// Get user from JWT token
	claims, err := middleware.GetUserClaims(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get room ID from URL
	vars := mux.Vars(r)
	roomIDStr := vars["room_id"]
	if roomIDStr == "" {
		http.Error(w, "Room ID is required", http.StatusBadRequest)
		return
	}

	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	// Verify user is a member of the room
	room, err := h.roomService.GetRoomByID(r.Context(), roomID)
	if err != nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	// Check if user is a member (room creator is automatically a member)
	// Get room members to check membership
	members, err := h.roomService.GetRoomMembers(r.Context(), roomID)
	if err != nil {
		http.Error(w, "Error checking room membership", http.StatusInternalServerError)
		return
	}

	isMember := room.CreatedBy == claims.UserID
	for _, member := range members {
		if member.UserID == claims.UserID {
			isMember = true
			break
		}
	}

	if !isMember {
		http.Error(w, "You are not a member of this room", http.StatusForbidden)
		return
	}

	// Upgrade connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	client := &Client{
		hub:    h.hub,
		conn:   conn,
		send:   make(chan []byte, 256),
		userID: claims.UserID.String(),
		roomID: roomID.String(),
	}

	h.hub.register <- client

	// Start goroutines
	go client.WritePump()
	client.ReadPump()
}

// processIncomingMessages processes messages from the hub broadcast channel
// This intercepts messages before they're sent to clients, saves them to DB, then rebroadcasts
