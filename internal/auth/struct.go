package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// model Access Token claims
type CustomClaims struct {
	UserID uint   `json:"user_id`
	Role   string `json:"role`
	jwt.RegisteredClaims
}

// model Refresh Token claims
type RefreshSession struct {
	Token     string
	UserID    uint
	ExpiresAt time.Time
}

// TokenPair - пара токенов - доступа и обновления
type TokenPair struct {
	AccessToken  string
	RefreshToken string
}
