package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GavinHemsada/go-backend/internal/middleware"
	"github.com/GavinHemsada/go-backend/internal/services"
	"github.com/GavinHemsada/go-backend/pkg/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type RoomHandler struct {
	roomService *services.RoomService
}

func NewRoomHandler(roomService *services.RoomService) *RoomHandler {
	return &RoomHandler{
		roomService: roomService,
	}
}

type CreateRoomRequest struct {
	Name     string `json:"name"`
	RoomType string `json:"room_type"`
}

// CreateRoom handles room creation
func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	// Get user from JWT
	claims, err := middleware.GetUserClaims(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req CreateRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	room, err := h.roomService.CreateRoom(r.Context(), req.Name, req.RoomType, claims.UserID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, room)
}

// GetRoomByID handles getting a room by ID
func (h *RoomHandler) GetRoomByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID, err := uuid.Parse(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid room ID")
		return
	}

	room, err := h.roomService.GetRoomByID(r.Context(), roomID)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, room)
}

// GetAllRooms handles getting all rooms
func (h *RoomHandler) GetAllRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.roomService.GetAllRooms(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, rooms)
}

// GetUserRooms handles getting all rooms for the current user
func (h *RoomHandler) GetUserRooms(w http.ResponseWriter, r *http.Request) {
	// Get user from JWT
	claims, err := middleware.GetUserClaims(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	rooms, err := h.roomService.GetUserRooms(r.Context(), claims.UserID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, rooms)
}

// DeleteRoom handles room deletion
func (h *RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	// Get user from JWT
	claims, err := middleware.GetUserClaims(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	roomID, err := uuid.Parse(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid room ID")
		return
	}

	err = h.roomService.DeleteRoom(r.Context(), roomID, claims.UserID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Room deleted successfully"})
}

// JoinRoom handles joining a room
func (h *RoomHandler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	// Get user from JWT
	claims, err := middleware.GetUserClaims(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	roomID, err := uuid.Parse(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid room ID")
		return
	}

	err = h.roomService.JoinRoom(r.Context(), roomID, claims.UserID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Joined room successfully"})
}

// LeaveRoom handles leaving a room
func (h *RoomHandler) LeaveRoom(w http.ResponseWriter, r *http.Request) {
	// Get user from JWT
	claims, err := middleware.GetUserClaims(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	roomID, err := uuid.Parse(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid room ID")
		return
	}

	err = h.roomService.LeaveRoom(r.Context(), roomID, claims.UserID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Left room successfully"})
}

// GetRoomMembers handles getting all members of a room
func (h *RoomHandler) GetRoomMembers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID, err := uuid.Parse(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid room ID")
		return
	}

	members, err := h.roomService.GetRoomMembers(r.Context(), roomID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, members)
}
