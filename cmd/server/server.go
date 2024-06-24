package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/WasathTheekshana/golang-project-structure/interfaces"
)

type GinServer interface {
	Start(ctx context.Context, httpAddress string) error
	Shutdown(ctx context.Context) error
	RegisterGroupRoute(path string, routes []interfaces.IRouteDefinition, middlewares ...gin.HandlerFunc)
	RegisterRoute(method, path string, handler gin.HandlerFunc)
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

// Method to register a single route
func (gs *ginServer) RegisterRoute(method, path string, handler gin.HandlerFunc) {
	switch method {
	case "GET":
		gs.engine.GET(path, handler)
	case "POST":
		gs.engine.POST(path, handler)
	case "PUT":
		gs.engine.PUT(path, handler)
	case "DELETE":
		gs.engine.DELETE(path, handler)
	case "PATCH":
		gs.engine.PATCH(path, handler)
	default:
		log.Errorf("Invalid HTTP method provided")
	}
}

// Register group of routes
func (gs *ginServer) RegisterGroupRoute(path string, routes []interfaces.IRouteDefinition, middlewares ...gin.HandlerFunc) {
	group := gs.engine.Group("api/v1" + path)
	group.Use(middlewares...)
	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Path, route.Handler)
		case "POST":
			group.POST(route.Path, route.Handler)
		case "PUT":
			group.PUT(route.Path, route.Handler)
		case "DELETE":
			group.DELETE(route.Path, route.Handler)
		case "PATCH":
			group.PATCH(route.Path, route.Handler)
		default:
			log.Errorf("Invalid HTTP method provided")
		}
	}
}
