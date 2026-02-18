package services

import (
	"context"
	"errors"

	"github.com/GavinHemsada/go-backend/internal/models"
	repository "github.com/GavinHemsada/go-backend/internal/repositories"
	"github.com/google/uuid"
)

type MessageService struct {
	messageRepo *repository.MessageRepository
	roomRepo    *repository.RoomRepository
}

func NewMessageService(messageRepo *repository.MessageRepository, roomRepo *repository.RoomRepository) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
		roomRepo:    roomRepo,
	}
}

// CreateMessage creates a new message in a room
func (s *MessageService) CreateMessage(ctx context.Context, roomID, userID uuid.UUID, content, messageType string) (*models.Message, error) {
	if content == "" {
		return nil, errors.New("message content is required")
	}

	// Check if user is a member of the room
	isMember, err := s.roomRepo.IsMember(ctx, roomID, userID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, errors.New("user is not a member of this room")
	}

	if messageType == "" {
		messageType = "text" // Default message type
	}

	message := &models.Message{
		RoomID:      roomID,
		UserID:      userID,
		Content:     content,
		MessageType: messageType,
	}

	err = s.messageRepo.Create(ctx, message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

// GetMessagesByRoom retrieves messages from a room with pagination
func (s *MessageService) GetMessagesByRoom(ctx context.Context, roomID uuid.UUID, limit, offset int) ([]models.Message, error) {
	if limit <= 0 {
		limit = 50 // Default limit
	}
	if limit > 100 {
		limit = 100 // Max limit
	}
	if offset < 0 {
		offset = 0
	}

	return s.messageRepo.GetByRoom(ctx, roomID, limit, offset)
}
