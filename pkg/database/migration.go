package database

import (
	"fmt"
	"github.com/pkg/errors"
	"real-estate-management/pkg/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

const migrationsPath = "deployment/migration"

// MigrateDB applies database migrations
func MigrateDB(dbConfig config.DatabaseConfig) {
	log.Info().Msg("Starting database migration...")

	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationsPath),
		dbConfig.GetDSN(),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create migration instance")
	}

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Warn().Msg("No new migrations to apply")
		} else {
			log.Fatal().Err(err).Msg("Migration failed")
		}
	} else {
		log.Info().Msg("Database migration completed successfully")
	}
}

// RollbackDB rolls back the last migration step
func RollbackDB(dbConfig config.DatabaseConfig) {
	log.Warn().Msg("Rolling back the last migration...")

	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationsPath),
		dbConfig.GetDSN(),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create migration instance")
	}

	err = m.Steps(-1)
	if err != nil {
		log.Fatal().Err(err).Msg("Rollback failed")
	}

	log.Info().Msg("Database rollback successful")
}
