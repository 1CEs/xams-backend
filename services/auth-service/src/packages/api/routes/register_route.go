package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xams-backend/services/auth-service/src/internal/users/repository"
	"github.com/xams-backend/services/auth-service/src/internal/users/usecase"
	"github.com/xams-backend/services/auth-service/src/packages/api/controllers"
	"github.com/xams-backend/services/auth-service/src/packages/auth"
	"gorm.io/gorm"
)

func RegisterRoute(db *gorm.DB, router *gin.RouterGroup) {
	userUsecase := usecase.NewUserUsecase(repository.NewUserRepository(db))
	authentication := auth.NewAuthentication(userUsecase)
	
	registerController := controllers.RegisterController{
		Auth: authentication,
	}

	register_path := os.Getenv("REGISTER_ROUTE_PATH")
	router.GET(register_path, registerController.Register)
}