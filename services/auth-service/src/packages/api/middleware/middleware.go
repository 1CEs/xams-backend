package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xams-backend/services/auth-service/src/packages/utils/middleware"
)

type Middleware struct{}

func (m *Middleware) IsAuthorized(context *gin.Context) {
	cookieName := os.Getenv("COOKIE_JWT_TOKEN")
	token, err := context.Cookie(cookieName)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	mu := middleware.MiddlewareUtils{}
	claims, err := mu.ParsedJWT(token)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}

	log.Println(claims)

	context.Next()
}