package services

import (
	"context"
	"errors"

	"github.com/GavinHemsada/go-backend/internal/models"
	"github.com/GavinHemsada/go-backend/internal/dtos"
	repository "github.com/GavinHemsada/go-backend/internal/repositories"
	"github.com/GavinHemsada/go-backend/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

func NewUserService(userRepo *repository.UserRepository, jwtSecret string) *UserService {
	return &UserService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// Register creates a new user and returns the user with a JWT token
func (s *UserService) Register(ctx context.Context, username, email, password string) (*dtos.AuthResponse, error) {
	// Validate input
	if username == "" || email == "" || password == "" {
		return nil, errors.New("username, email, and password are required")
	}

	if len(password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	// Create user via repository
	user, err := s.userRepo.Register(ctx, username, email, password)
	if err != nil {
		return nil, err
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &dtos.AuthResponse{
		User:  user,
		Token: token,
	}, nil
}

// Login authenticates a user and returns the user with a JWT token
func (s *UserService) Login(ctx context.Context, identifier, password string) (*dtos.AuthResponse, error) {
	// Validate input
	if identifier == "" || password == "" {
		return nil, errors.New("identifier and password are required")
	}

	// Authenticate user via repository
	user, err := s.userRepo.Login(ctx, identifier, password)
	if err != nil {
		return nil, err
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &dtos.AuthResponse{
		User:  user,
		Token: token,
	}, nil
}

// GetByID retrieves a user by their ID
func (s *UserService) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

// GetAll retrieves all users
func (s *UserService) GetAll(ctx context.Context) ([]models.User, error) {
	return s.userRepo.GetAll(ctx)
}

// ValidateToken validates a JWT token and returns the claims
func (s *UserService) ValidateToken(tokenString string) (*utils.Claims, error) {
	claims := &utils.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}