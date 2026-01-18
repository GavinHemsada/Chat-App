package dtos

import "github.com/GavinHemsada/go-backend/internal/models"

type AuthResponse struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}