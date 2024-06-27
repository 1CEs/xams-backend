package middleware

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/xams-backend/services/auth-service/src/internal/models"
)

type MiddlewareUtils struct{}

func (mu *MiddlewareUtils) ParsedJWT(tokenString string) (*models.UserClaims, error) {
	claims := &models.UserClaims{}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET environment variable not set")
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}