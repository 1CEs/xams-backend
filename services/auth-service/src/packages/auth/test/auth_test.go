package test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/xams-backend/services/auth-service/src/internal/models"
	"github.com/xams-backend/services/auth-service/src/internal/users/repository"
	"github.com/xams-backend/services/auth-service/src/internal/users/usecase"
	"github.com/xams-backend/services/auth-service/src/packages/auth"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	dsn_test := "root:123@tcp(localhost:3306)/xams?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn_test), &gorm.Config{})
	return db
}

func Test_Register(t *testing.T) {
	db := setupTestDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	authen := auth.NewAuthentication(userUsecase)
	user, _, err := authen.Register(&models.User{
		UserID: "1111111111-8",
		Password: "helloja232",
		Email: "ic21312e@gmail.com",
		Prename: "nai",
		FirstName: "exam",
		LastName: "ple",
		BranchID: 1,
		Role: "teacher",
	})
	userJson, _ := json.MarshalIndent(user, "", "		")
	log.Println(string(userJson))
	log.Println(err)
}

func Test_Login(t *testing.T) {
	db := setupTestDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	authen := auth.NewAuthentication(userUsecase)

	user, _, err := authen.Login("ice@gmail.com", "helloja")
	userJson, _ := json.MarshalIndent(user, "", "		")
	log.Println(string(userJson))
	log.Println(err)
}