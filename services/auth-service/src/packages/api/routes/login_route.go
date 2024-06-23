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

func LoginRoute(db *gorm.DB, router *gin.RouterGroup) {
	userUsecase := usecase.NewUserUsecase(repository.NewUserRepository(db))
	authentication := auth.NewAuthentication(userUsecase)
	
	loginController := controllers.LoginController{
		Auth: authentication,
	}

	login_path := os.Getenv("LOGIN_ROUTE_PATH")
	router.GET(login_path, loginController.Login)
}