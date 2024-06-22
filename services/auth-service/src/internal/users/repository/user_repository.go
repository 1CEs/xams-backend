package repository

import (
	"github.com/xams-backend/services/auth-service/src/internal/models"
	"gorm.io/gorm"
)

type (
	UserRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetUser(id string) (*models.UserResponse, error) {
	var user models.User

	if err := repo.db.First(&user, "user_id = ?", id).Error; err != nil {
		return nil, err
	}

	userResponse := models.UserResponse{
		UserID: user.UserID,
		Email: user.Email,
		Prename: user.Prename,
		FirstName: user.FirstName,
		LastName: user.LastName,
		BranchID: user.BranchID,
		Role: user.Role,
	}

	return &userResponse, nil
}

func (repo *UserRepository) CreateUser(user *models.User) error {

	if err := repo.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) UpdateUser(user *models.User) error {
	if err := repo.db.Model(&models.User{}).Where("user_id = ?", user.UserID).Updates(user).Error; err != nil {
		return err
	}

	return nil
}