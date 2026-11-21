package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenPair(userID uint, role string) (*TokenPair, error) {

	// секретный ключ
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// Access Token ->
	accessClaims := CustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessStr, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("Ошибка подписи access токена^: %w", err)
	}
	// <- Access Token

	// Refresh Token ->
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return nil, fmt.Errorf("Ошибка получения случайных криптографических значений: %w", err)
	}
	refreshStr := hex.EncodeToString(b)
	// <- Refresh Token

	return &TokenPair{
		AccessToken:  accessStr,
		RefreshToken: refreshStr,
	}, nil
}
