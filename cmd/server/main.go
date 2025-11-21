package main

import (
	"log"

	"fitness-app/internal/db"
	"fitness-app/internal/domain/models"
	"fitness-app/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment")
	}

	// Connect DB
	db.Connect()
	db.AutoMigrate(&models.User{}, &models.Workout{}, &models.Session{})

	// Start server
	if err := server.Run(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
