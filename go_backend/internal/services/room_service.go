package services

import (
	"context"
	"errors"

	"github.com/GavinHemsada/go-backend/internal/models"
	repository "github.com/GavinHemsada/go-backend/internal/repositories"
	"github.com/google/uuid"
)

type RoomService struct {
	roomRepo *repository.RoomRepository
}

func NewRoomService(roomRepo *repository.RoomRepository) *RoomService {
	return &RoomService{
		roomRepo: roomRepo,
	}
}

// CreateRoom creates a new room
func (s *RoomService) CreateRoom(ctx context.Context, name, roomType string, createdBy uuid.UUID) (*models.Room, error) {
	if name == "" {
		return nil, errors.New("room name is required")
	}

	if roomType == "" {
		roomType = "public" // Default room type
	}

	room := &models.Room{
		Name:      name,
		RoomType:  roomType,
		CreatedBy: createdBy,
	}

	err := s.roomRepo.Create(ctx, room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

// GetRoomByID retrieves a room by ID
func (s *RoomService) GetRoomByID(ctx context.Context, roomID uuid.UUID) (*models.Room, error) {
	return s.roomRepo.GetByID(ctx, roomID)
}

// GetAllRooms retrieves all rooms
func (s *RoomService) GetAllRooms(ctx context.Context) ([]models.Room, error) {
	return s.roomRepo.GetAll(ctx)
}

// GetUserRooms retrieves all rooms a user is a member of
func (s *RoomService) GetUserRooms(ctx context.Context, userID uuid.UUID) ([]models.Room, error) {
	return s.roomRepo.GetUserRooms(ctx, userID)
}

// DeleteRoom deletes a room (only creator can delete)
func (s *RoomService) DeleteRoom(ctx context.Context, roomID, userID uuid.UUID) error {
	return s.roomRepo.Delete(ctx, roomID, userID)
}

// JoinRoom adds a user to a room
func (s *RoomService) JoinRoom(ctx context.Context, roomID, userID uuid.UUID) error {
	// Check if room exists
	_, err := s.roomRepo.GetByID(ctx, roomID)
	if err != nil {
		return err
	}

	return s.roomRepo.AddMember(ctx, roomID, userID)
}

// LeaveRoom removes a user from a room
func (s *RoomService) LeaveRoom(ctx context.Context, roomID, userID uuid.UUID) error {
	return s.roomRepo.RemoveMember(ctx, roomID, userID)
}

// GetRoomMembers retrieves all members of a room
func (s *RoomService) GetRoomMembers(ctx context.Context, roomID uuid.UUID) ([]models.RoomMember, error) {
	return s.roomRepo.GetMembers(ctx, roomID)
}
