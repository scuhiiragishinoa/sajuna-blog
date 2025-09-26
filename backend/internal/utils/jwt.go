package utils

import (
	"errors"
	"sajuna-blog-backend/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ValidateToken 验证JWT token
func ValidateToken(tokenString string, cfg *config.Config) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
