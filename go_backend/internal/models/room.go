package models

import (
    "time"
    "github.com/google/uuid"
)

type Room struct {
    ID        uuid.UUID `json:"id" db:"id"`
    Name      string    `json:"name" db:"name"`
    RoomType  string    `json:"room_type" db:"room_type"`
    CreatedBy uuid.UUID `json:"created_by" db:"created_by"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
}