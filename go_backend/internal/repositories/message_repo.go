package repository

import (
    "context"
    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
    "github.com/GavinHemsada/go-backend/internal/models"
)

type MessageRepository struct {
    db *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) *MessageRepository {
    return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(ctx context.Context, msg *models.Message) error {
    query := `
        INSERT INTO messages (room_id, user_id, content, message_type)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
    `
    return r.db.QueryRowContext(
        ctx, query,
        msg.RoomID, msg.UserID, msg.Content, msg.MessageType,
    ).Scan(&msg.ID, &msg.CreatedAt)
}

func (r *MessageRepository) GetByRoom(ctx context.Context, roomID uuid.UUID, limit, offset int) ([]models.Message, error) {
    query := `
        SELECT m.id, m.room_id, m.user_id, m.content, m.message_type, m.created_at, u.username
        FROM messages m
        JOIN users u ON m.user_id = u.id
        WHERE m.room_id = $1 AND m.is_deleted = false
        ORDER BY m.created_at DESC
        LIMIT $2 OFFSET $3
    `
    
    var messages []models.Message
    err := r.db.SelectContext(ctx, &messages, query, roomID, limit, offset)
    return messages, err
}