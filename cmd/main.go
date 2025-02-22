package main

import (
	"fmt"

	"real-estate-management/internal/api/router"
	"real-estate-management/pkg/config"
	"real-estate-management/pkg/database"

	"github.com/rs/zerolog/log"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Create the database if it does not exist
	database.CreateDB(cfg.Database)

	// Run database migrations using Golang Migrate
	database.MigrateDB(cfg.Database)

	// Connect to database
	db := database.ConnectDB(cfg.Database)

	// Initialize and start the server
	app := router.SetupRoutes(db)

	port := fmt.Sprintf(":%d", cfg.App.Port)
	log.Info().Msgf("Starting server on port %s...", port)

	if err := app.Listen(port); err != nil {
		log.Fatal().Err(err).Msg("Error starting server")
	}
}
