package database

import (
	"context"
	"real-estate-management/pkg/config"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GormLogger implements GORM's logger.Interface using zerolog
type GormLogger struct {
	SlowThreshold time.Duration
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	log.Info().Msgf(msg, data...)
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	log.Warn().Msgf(msg, data...)
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	log.Error().Msgf(msg, data...)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil {
		log.Error().
			Err(err).
			Str("sql", sql).
			Int64("rows", rows).
			Dur("elapsed", elapsed).
			Msg("SQL Error")
		return
	}

	if elapsed > l.SlowThreshold {
		log.Warn().
			Str("sql", sql).
			Int64("rows", rows).
			Dur("elapsed", elapsed).
			Msg("Slow SQL Query")
		return
	}

	log.Debug().
		Str("sql", sql).
		Int64("rows", rows).
		Dur("elapsed", elapsed).
		Msg("SQL Query")
}

// ConnectDB establishes database connection using global logger
func ConnectDB(dbConfig config.DatabaseConfig) *gorm.DB {
	// Create GORM logger
	gormLogger := &GormLogger{
		SlowThreshold: time.Second,
	}

	// Configure GORM
	gormConfig := &gorm.Config{
		Logger: gormLogger,
	}

	// Attempt database connection
	db, err := gorm.Open(postgres.Open(dbConfig.GetDSN()), gormConfig)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("host", dbConfig.Host).
			Int("port", dbConfig.Port).
			Str("database", dbConfig.Name).
			Msg("Failed to connect to database")
	}

	log.Info().
		Str("host", dbConfig.Host).
		Int("port", dbConfig.Port).
		Str("database", dbConfig.Name).
		Msg("Connected to database successfully")

	return db
}
