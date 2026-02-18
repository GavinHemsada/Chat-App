package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/GavinHemsada/go-backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RoomRepository struct {
	db *sqlx.DB
}

func NewRoomRepository(db *sqlx.DB) *RoomRepository {
	return &RoomRepository{db: db}
}

// Create creates a new room
func (r *RoomRepository) Create(ctx context.Context, room *models.Room) error {
	room.ID = uuid.New()
	query := `
		INSERT INTO rooms (id, name, room_type, created_by)
		VALUES ($1, $2, $3, $4)
		RETURNING created_at
	`
	err := r.db.QueryRowContext(
		ctx, query,
		room.ID, room.Name, room.RoomType, room.CreatedBy,
	).Scan(&room.CreatedAt)

	if err != nil {
		return err
	}

	// Automatically add creator as a member
	return r.AddMember(ctx, room.ID, room.CreatedBy)
}

// GetByID retrieves a room by its ID
func (r *RoomRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Room, error) {
	var room models.Room
	query := `
		SELECT id, name, room_type, created_by, created_at
		FROM rooms
		WHERE id = $1
	`
	err := r.db.GetContext(ctx, &room, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("room not found")
		}
		return nil, err
	}
	return &room, nil
}

// GetAll retrieves all rooms
func (r *RoomRepository) GetAll(ctx context.Context) ([]models.Room, error) {
	var rooms []models.Room
	query := `
		SELECT id, name, room_type, created_by, created_at
		FROM rooms
		ORDER BY created_at DESC
	`
	err := r.db.SelectContext(ctx, &rooms, query)
	return rooms, err
}

// GetUserRooms retrieves all rooms a user is a member of
func (r *RoomRepository) GetUserRooms(ctx context.Context, userID uuid.UUID) ([]models.Room, error) {
	var rooms []models.Room
	query := `
		SELECT r.id, r.name, r.room_type, r.created_by, r.created_at
		FROM rooms r
		INNER JOIN room_members rm ON r.id = rm.room_id
		WHERE rm.user_id = $1
		ORDER BY r.created_at DESC
	`
	err := r.db.SelectContext(ctx, &rooms, query, userID)
	return rooms, err
}

// Delete deletes a room (only if user is the creator)
func (r *RoomRepository) Delete(ctx context.Context, roomID, userID uuid.UUID) error {
	// Check if user is the creator
	room, err := r.GetByID(ctx, roomID)
	if err != nil {
		return err
	}

	if room.CreatedBy != userID {
		return errors.New("only room creator can delete the room")
	}

	query := `DELETE FROM rooms WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, roomID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("room not found")
	}

	return nil
}

// AddMember adds a user to a room
func (r *RoomRepository) AddMember(ctx context.Context, roomID, userID uuid.UUID) error {
	query := `
		INSERT INTO room_members (room_id, user_id)
		VALUES ($1, $2)
		ON CONFLICT (room_id, user_id) DO NOTHING
	`
	_, err := r.db.ExecContext(ctx, query, roomID, userID)
	return err
}

// RemoveMember removes a user from a room
func (r *RoomRepository) RemoveMember(ctx context.Context, roomID, userID uuid.UUID) error {
	query := `DELETE FROM room_members WHERE room_id = $1 AND user_id = $2`
	result, err := r.db.ExecContext(ctx, query, roomID, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user is not a member of this room")
	}

	return nil
}

// GetMembers retrieves all members of a room
func (r *RoomRepository) GetMembers(ctx context.Context, roomID uuid.UUID) ([]models.RoomMember, error) {
	var members []models.RoomMember
	query := `
		SELECT room_id, user_id, joined_at
		FROM room_members
		WHERE room_id = $1
		ORDER BY joined_at ASC
	`
	err := r.db.SelectContext(ctx, &members, query, roomID)
	return members, err
}

// IsMember checks if a user is a member of a room
func (r *RoomRepository) IsMember(ctx context.Context, roomID, userID uuid.UUID) (bool, error) {
	var count int
	query := `
		SELECT COUNT(*)
		FROM room_members
		WHERE room_id = $1 AND user_id = $2
	`
	err := r.db.GetContext(ctx, &count, query, roomID, userID)
	return count > 0, err
}