package auth

import (
	"github.com/xams-backend/services/auth-service/src/internal/models"
)

type (
	IAuthentication interface {
		Login(email string, password string) (*models.LogInResponse, string, error)
		Register(user *models.User) (*models.LogInResponse, string, error)
		generateJWT(user *models.User) (string, error)
	}
)