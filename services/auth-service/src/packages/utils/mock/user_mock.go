package mock

import (
	"github.com/bxcodec/faker/v3"
	"github.com/xams-backend/services/auth-service/src/internal/models"
)

func GenerateMockUser() *models.User {
	return &models.User{
		UserID:    faker.CCNumber(),
		Password:  faker.Password(),
		Email:     faker.Email(),
		Prename:   "นาย",
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		BranchID:  1,
		Role:      models.Student,
	}
}