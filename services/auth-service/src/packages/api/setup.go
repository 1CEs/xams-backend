package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xams-backend/services/auth-service/src/packages/api/middleware"
	"github.com/xams-backend/services/auth-service/src/packages/api/routes"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, gin *gin.Engine) {
	// For authentication route
	route := gin.Group("/auth")
	routes.LoginRoute(db, route)
	routes.RegisterRoute(db, route)

	// For user route
	middleware := middleware.Middleware{}

	route = gin.Group("")
	route.Use(middleware.IsAuthorized)
	routes.UserRoute(db, route)

}