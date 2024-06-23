package repository

import "github.com/xams-backend/services/auth-service/src/internal/models"

type (
	IUserRepository interface {
		GetUser(id string) (*models.User, error)
		GetUserByEmail(email string) (*models.User, error)
		UpdateUser(user *models.User) error
		CreateUser(user *models.User) error
		IsUserAlreadyExists(id string) error
	}
)