package middleware

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/xams-backend/services/auth-service/src/internal/users/repository"
	"github.com/xams-backend/services/auth-service/src/internal/users/usecase"
	"github.com/xams-backend/services/auth-service/src/packages/database"
	"github.com/xams-backend/services/auth-service/src/packages/utils/middleware"
)

type Middleware struct{}

func (m *Middleware) IsAuthorized(context *gin.Context) {
	cookieName := os.Getenv("COOKIE_JWT_TOKEN")
	token, err := context.Cookie(cookieName)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		context.Abort()
		return
	}

	mu := middleware.MiddlewareUtils{}
	claims, err := mu.ParsedJWT(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		context.Abort()
		return
	}
	log.Println(claims.UserID)	
	db := database.NewDatabase()
	connect, err := db.Connect()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	ur := repository.NewUserRepository(connect)
	uc := usecase.NewUserUsecase(ur)

	if err := uc.IsUserAlreadyExists(claims.UserID); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Your credentials is not in our service."})
		context.Abort()
		return
	}
	
	context.Next()
}