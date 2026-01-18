package models

import (
    "time"
    "github.com/google/uuid"
)

type RoomMember struct {
    RoomID uuid.UUID `json:"room_id" db:"room_id"`
    UserID uuid.UUID `json:"user_id" db:"user_id"`
    JoinedAt time.Time `json:"joined_at" db:"joined_at"`
}