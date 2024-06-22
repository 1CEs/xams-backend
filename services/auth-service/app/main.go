package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/xams-backend/services/auth-service/src/packages/database"
)

func main() {
	godotenv.Load() //Load all environments

	db := database.NewDatabase()
	_, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if mode := os.Getenv("MODE"); mode == "production" {
		err = db.AutoMigration()
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
	}
	
}
