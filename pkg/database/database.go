package database

import (
	"fmt"
	"log"

	"real-estate-management/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dbConfig config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database successfully")

	return db
}
