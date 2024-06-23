package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnvionment() {
	if err := godotenv.Load("app.env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func SetupDatabse() (*gorm.DB, error) {
	dsn := os.Getenv(DatbaseURI)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %s", err)
	}

	return db, nil
}
