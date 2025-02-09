package config

import (
	"log"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Name     string `env:"DB_NAME,required"`
	SSLMode  string `env:"DB_SSL_MODE" envDefault:"disable"`
}

type AppConfig struct {
	Port     int    `env:"SERVER_PORT" envDefault:"8080"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

func LoadConfig() Config {
	var c Config

	// Load .env file only in non-production environments
	if os.Getenv("ENV") != "production" {
		_ = godotenv.Load()
	}

	// Parse environment variables
	if err := env.Parse(&c); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	return c
}
