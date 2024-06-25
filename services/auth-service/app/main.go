package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/xams-backend/services/auth-service/src/packages/api"
	"github.com/xams-backend/services/auth-service/src/packages/database"
)

func main() {
	godotenv.Load() // Load all environments

	// Initialize database
	db := database.NewDatabase()
	connect, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if mode := os.Getenv("MODE"); mode == "production" {
		err = db.AutoMigration()
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
	}
	router := gin.Default()
	api.SetupRoutes(connect, router)
	router.Run(":8000")
	
}
