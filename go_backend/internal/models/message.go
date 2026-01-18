package models

import (
    "time"
    "github.com/google/uuid"
)

type Message struct {
    ID          uuid.UUID `json:"id" db:"id"`
    RoomID      uuid.UUID `json:"room_id" db:"room_id"`
    UserID      uuid.UUID `json:"user_id" db:"user_id"`
    Content     string    `json:"content" db:"content"`
    MessageType string    `json:"message_type" db:"message_type"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    Username    string    `json:"username,omitempty"` // For display
}