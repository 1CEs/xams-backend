package repository

import "github.com/xams-backend/services/auth-service/src/internal/models"

type (
	IUserRepository interface {
		GetUser(id string) (*models.UserResponse, error)
		UpdateUser(user *models.User) error
		CreateUser(user *models.User) error
	}
)