package websocket

import (
    "encoding/json"
    "log"
    "net/http"
    
    "github.com/gorilla/websocket"
    "github.com/google/uuid"
    "github.com/GavinHemsada/go-backend/internal/models"
    "github.com/GavinHemsada/go-backend/internal/repositories"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true // Configure properly for production
    },
}

type Handler struct {
    hub      *Hub
    msgRepo  *repository.MessageRepository
}

func NewHandler(hub *Hub, msgRepo *repository.MessageRepository) *Handler {
    return &Handler{
        hub:     hub,
        msgRepo: msgRepo,
    }
}

func (h *Handler) ServeWS(w http.ResponseWriter, r *http.Request) {
    // Extract user and room from query params or JWT
    userID := r.URL.Query().Get("user_id")
    roomID := r.URL.Query().Get("room_id")
    
    if userID == "" || roomID == "" {
        http.Error(w, "Missing user_id or room_id", http.StatusBadRequest)
        return
    }
    
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("WebSocket upgrade error: %v", err)
        return
    }

    client := &Client{
        hub:    h.hub,
        conn:   conn,
        send:   make(chan []byte, 256),
        userID: userID,
        roomID: roomID,
    }
    
    h.hub.register <- client
    
    // Start goroutines
    go client.WritePump()
    go h.handleMessages(client)
    client.ReadPump()
}

func (h *Handler) handleMessages(client *Client) {
    for message := range client.send {
        var wsMsg models.WSMessage
        if err := json.Unmarshal(message, &wsMsg); err != nil {
            continue
        }
        
        // Save message to database
        if wsMsg.Type == "message" {
            msg := &models.Message{
                RoomID:      uuid.MustParse(wsMsg.RoomID),
                UserID:      uuid.MustParse(client.userID),
                Content:     wsMsg.Payload.(string),
                MessageType: "text",
            }
            
            if err := h.msgRepo.Create(r.Context(), msg); err != nil {
                log.Printf("Error saving message: %v", err)
            }
        }
    }
}