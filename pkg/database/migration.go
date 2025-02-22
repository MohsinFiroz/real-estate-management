package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"real-estate-management/pkg/config"
)

// CreateDB creates the database if it does not exist
func CreateDB(dbConfig config.DatabaseConfig) {
	// Connect to the PostgreSQL server (without specifying a database)
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password)

	// Open connection to the server
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to PostgreSQL server")
	}

	// Create the database if it doesn't exist
	err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbConfig.Name)).Error
	if err != nil {
		if err.Error() != "pq: database \"real-estate-management\" already exists" {
			log.Fatal().Err(err).Msg("failed to create database")
		}
		log.Warn().Msg("Database already exists")
	} else {
		log.Info().Msg("Database created successfully")
	}

}

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
