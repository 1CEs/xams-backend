package auth

import (
	"github.com/xams-backend/services/auth-service/src/internal/models"
)

type (
	IAuthentication interface {
		Login(id string, password string) (*models.UserResponse, error)
		Register(user *models.User) (*models.UserResponse, error)
		generateJWT(user *models.User) (string, error)
	}
)