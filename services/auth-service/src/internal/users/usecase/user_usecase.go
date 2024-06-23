package usecase

import (
	"github.com/xams-backend/services/auth-service/src/internal/models"
	"github.com/xams-backend/services/auth-service/src/internal/users/repository"
)

type (
	UserUsecase struct {
		repo repository.IUserRepository
	}
)

func NewUserUsecase(repo repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (usecase *UserUsecase ) GetUser(id string) (*models.User, error) {
	return usecase.repo.GetUser(id)
}

func (usecase *UserUsecase ) CreateUser(user *models.User) error {
	return usecase.repo.CreateUser(user)
}

func (usecase *UserUsecase ) UpdateUser(user *models.User) error {
	return usecase.repo.UpdateUser(user)
}

func (usecase *UserUsecase ) IsUserAlreadyExists(id string) error {
	return usecase.repo.IsUserAlreadyExists(id)
}