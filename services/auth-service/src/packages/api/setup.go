package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xams-backend/services/auth-service/src/packages/api/middleware"
	"github.com/xams-backend/services/auth-service/src/packages/api/routes"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, gin *gin.Engine) {
	middleware := middleware.Middleware{}

	// For authentication route
	auth_route := gin.Group("/auth")
	routes.LoginRoute(db, auth_route)
	routes.RegisterRoute(db, auth_route)

	// For user route
	user_route := gin.Group("")
	user_route.Use(middleware.IsAuthorized)
	routes.UserRoute(db, user_route)
	
}
