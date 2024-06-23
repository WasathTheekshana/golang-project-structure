package cmd

import (
	"context"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/WasathTheekshana/golang-project-structure/cmd/server"
	"github.com/WasathTheekshana/golang-project-structure/config"
)

func Execute() {
	builder := server.NewGinServerBuilder()
	server := builder.Build()

	ctx := context.Background()

	config.LoadEnvionment()
	_, err := config.SetupDatabse()
	if err != nil {
		log.Fatal("Error setting up the database %v", err)
	}

	go func() {
		if err := server.Start(ctx, os.Getenv(config.AppPort)); err != nil {
			log.Errorf("Error starting the server %v", err)
		}
	}()

	<-ctx.Done()
	log.Info("Server stopped")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Errorf("Error shutting donw the server %v", err)
	}
}
