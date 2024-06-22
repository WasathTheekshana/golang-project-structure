package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func LoadEnvionment() {
	if err := godotenv.Load("app.env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
