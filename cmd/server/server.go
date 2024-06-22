package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type GinServer interface {
	Start(ctx context.Context, httpAddress string) error
	Shutdown(ctx context.Context) error
}

type GinServerBuilder struct{}

// Gin server
type ginServer struct {
	engine *gin.Engine
	server *http.Server
}

func NewGinServerBuilder() *GinServerBuilder {
	return &GinServerBuilder{}
}

func (b *GinServerBuilder) Build() GinServer {
	engine := gin.Default()
	return &ginServer{engine: engine}
}

func (gs *ginServer) Start(ctx context.Context, httpAddress string) error {
	gs.server = &http.Server{
		Addr:    httpAddress,
		Handler: gs.engine,
	}

	go func() {
		if err := gs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n ", err)
		}
	}()

	log.Infof("Server is running on port %s", httpAddress)

	return nil
}

func (gs *ginServer) Shutdown(ctx context.Context) error {
	if err := gs.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %s", err)
	}

	log.Infof("Server is shutting down...")
	return nil
}
