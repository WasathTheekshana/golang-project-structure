package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/WasathTheekshana/golang-project-structure/internal/model"
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

	if err := RunAutoMigrate(db); err != nil {
		return nil, fmt.Errorf("error running auto migration: %s", err)
	}

	return db, nil
}

func RunAutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.User{},
	); err != nil {
		errorMessage := fmt.Sprintf("error runnign auto migration: %s", err)
		log.Error(errorMessage)
		return errors.New(errorMessage)
	}
	return nil
}
