package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/GavinHemsada/go-backend/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserClaimsKey contextKey = "userClaims"

// JWTMiddleware creates a middleware that validates JWT tokens
func JWTMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract token from Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.RespondWithError(w, http.StatusUnauthorized, "Authorization header is required")
				return
			}

			// Check if it starts with "Bearer "
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				utils.RespondWithError(w, http.StatusUnauthorized, "Invalid authorization header format. Expected: Bearer <token>")
				return
			}

			tokenString := parts[1]

			// Validate token
			claims := &utils.Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				// Validate signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("invalid signing method")
				}
				return []byte(jwtSecret), nil
			})

			if err != nil {
				utils.RespondWithError(w, http.StatusUnauthorized, "Invalid or expired token")
				return
			}

			if !token.Valid {
				utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
				return
			}

			// Add claims to request context
			ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserClaims extracts user claims from the request context
func GetUserClaims(r *http.Request) (*utils.Claims, error) {
	claims, ok := r.Context().Value(UserClaimsKey).(*utils.Claims)
	if !ok {
		return nil, errors.New("user claims not found in context")
	}
	return claims, nil
}
