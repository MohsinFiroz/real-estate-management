package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"real-estate-management/pkg/config"
)

// MigrateDB applies database migrations using Golang Migrate
func MigrateDB(dbConfig config.DatabaseConfig) {
	// Set up the database connection string for migrations
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	// Create the migration instance
	m, err := migrate.New(
		fmt.Sprintf("file://%s", "deployment/migrations"), // Path to your migration files
		dsn,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create migration instance")
	}

	// Run the migrations (apply up migrations)
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal().Err(err).Msg("Migration failed")
	} else {
		log.Info().Msg("Migrations applied successfully")
	}
}
