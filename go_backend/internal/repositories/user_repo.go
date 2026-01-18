package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/GavinHemsada/go-backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
    db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
    return &UserRepository{db: db}
}

// Register creates a new user with a hashed password
func (r *UserRepository) Register(ctx context.Context, username, email, password string) (*models.User, error) {
    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &models.User{
        ID:           uuid.New(),
        Username:     username,
        Email:        email,
        PasswordHash: string(hashedPassword),
    }

    query := `
        INSERT INTO users (id, username, email, password_hash)
        VALUES ($1, $2, $3, $4)
        RETURNING created_at
    `
    
    err = r.db.QueryRowContext(
        ctx, query,
        user.ID, user.Username, user.Email, user.PasswordHash,
    ).Scan(&user.CreatedAt)
    
    if err != nil {
        return nil, err
    }

    return user, nil
}

// Login authenticates a user by email/username and password
func (r *UserRepository) Login(ctx context.Context, identifier, password string) (*models.User, error) {
    var user models.User
    
    // Try to find user by email or username
    query := `
        SELECT id, username, email, password_hash, created_at
        FROM users
        WHERE email = $1 OR username = $1
    `
    
    err := r.db.GetContext(ctx, &user, query, identifier)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("invalid credentials")
        }
        return nil, err
    }

    // Verify password
    err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
    if err != nil {
        return nil, errors.New("invalid credentials")
    }

    return &user, nil
}

// GetByID retrieves a user by their ID
func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
    var user models.User
    
    query := `
        SELECT id, username, email, password_hash, created_at
        FROM users
        WHERE id = $1
    `
    
    err := r.db.GetContext(ctx, &user, query, id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("user not found")
        }
        return nil, err
    }

    return &user, nil
}

// GetAll retrieves all users
func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
    var users []models.User
    
    query := `
        SELECT id, username, email, password_hash, created_at
        FROM users
        ORDER BY created_at DESC
    `
    
    err := r.db.SelectContext(ctx, &users, query)
    if err != nil {
        return nil, err
    }

    return users, nil
}
