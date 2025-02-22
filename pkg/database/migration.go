package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"real-estate-management/pkg/config"
)

// CreateDB creates the database if it does not exist
func CreateDB(dbConfig config.DatabaseConfig) {
	// Connect to the PostgreSQL server (without specifying a database)
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password)

	// Open connection to the server
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to PostgreSQL server")
	}
	defer db.Close()

	// Check if the target database already exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", dbConfig.Name).Scan(&exists)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to check database existence")
	}

	// If the database doesn't exist, create it
	if !exists {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbConfig.Name))
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "42P04" {
				log.Warn().Msg("Database already exists")
			} else {
				log.Fatal().Err(err).Msg("failed to create database")
			}
		} else {
			log.Info().Msg("Database created successfully")
		}
	} else {
		log.Info().Msg("Database already exists")
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
