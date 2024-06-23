package auth

import "github.com/xams-backend/services/auth-service/src/internal/models"

type (
	Authentication struct{}
)

func NewAuth() IAuthentication {
	return &Authentication{}
}

func (auth *Authentication) Login(id string, password string) (*models.UserResponse, error) {
	return &models.UserResponse{}, nil
}

func (auth *Authentication) Register(user *models.User) (*models.UserResponse, error) {
	return &models.UserResponse{}, nil
}