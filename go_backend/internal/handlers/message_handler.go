package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GavinHemsada/go-backend/internal/middleware"
	"github.com/GavinHemsada/go-backend/internal/services"
	"github.com/GavinHemsada/go-backend/pkg/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type MessageHandler struct {
	messageService *services.MessageService
}

func NewMessageHandler(messageService *services.MessageService) *MessageHandler {
	return &MessageHandler{
		messageService: messageService,
	}
}

type CreateMessageRequest struct {
	Content     string `json:"content"`
	MessageType string `json:"message_type"`
}

// CreateMessage handles message creation
func (h *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	// Get user from JWT
	claims, err := middleware.GetUserClaims(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	roomID, err := uuid.Parse(vars["room_id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid room ID")
		return
	}

	var req CreateMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	message, err := h.messageService.CreateMessage(r.Context(), roomID, claims.UserID, req.Content, req.MessageType)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, message)
}

// GetMessagesByRoom handles getting messages from a room
func (h *MessageHandler) GetMessagesByRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID, err := uuid.Parse(vars["room_id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid room ID")
		return
	}

	// Get pagination parameters
	limit := 50 // default
	offset := 0 // default

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil {
			limit = parsedLimit
		}
	}

	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		if parsedOffset, err := strconv.Atoi(offsetStr); err == nil {
			offset = parsedOffset
		}
	}

	messages, err := h.messageService.GetMessagesByRoom(r.Context(), roomID, limit, offset)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, messages)
}
