package routes

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/xams-backend/services/auth-service/src/internal/users/repository"
	"github.com/xams-backend/services/auth-service/src/internal/users/usecase"
	"github.com/xams-backend/services/auth-service/src/packages/api/controllers"
	"gorm.io/gorm"
)

func UserRoute(db *gorm.DB, router *gin.RouterGroup) {
	userUsecase := usecase.NewUserUsecase(repository.NewUserRepository(db))
	
	userController := controllers.UserController{
		Usecase: userUsecase,
	}

	user_path := os.Getenv("USER_ROUTE_PATH")
	router.GET(user_path + "/:id", userController.GetUser)
	router.GET(user_path + "/:id", userController.UpdateUser)
}