package cmd

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/WasathTheekshana/golang-project-structure/cmd/server"
	"github.com/WasathTheekshana/golang-project-structure/config"
)

func Execute() {
	builder := server.NewGinServerBuilder()
	server := builder.Build()

	ctx := context.Background()
	config.LoadEnvionment()

	go func() {
		if err := server.Start(ctx, os.Getenv(config.AppPort)); if err != nil {
			log.Errorf("Error starting the server %v", err)
		}
	}()

	<-ctx.Done()
	log.Info("Shutting down the server...")

	
}
