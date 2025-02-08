package main

import (
	"fmt"
	"log"

	"real-estate-management/internal/api/router"
	"real-estate-management/pkg/config"
	"real-estate-management/pkg/database"
)

func main() {
	// Load configuration
	config := config.LoadConfig()

	// Connect to database
	db := database.ConnectDB(config.Database)

	// Initialize and start the server
	app := router.SetupRoutes(db)

	port := fmt.Sprintf(":%d", config.App.Port)
	log.Printf("Starting server on port %s...", port)

	if err := app.Listen(port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
