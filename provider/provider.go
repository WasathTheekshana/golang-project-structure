package provider

import (
	"gorm.io/gorm"

	"github.com/WasathTheekshana/golang-project-structure/cmd/server"
	"github.com/WasathTheekshana/golang-project-structure/internal/handler"
	"github.com/WasathTheekshana/golang-project-structure/internal/repository"
	"github.com/WasathTheekshana/golang-project-structure/internal/routes"
	"github.com/WasathTheekshana/golang-project-structure/internal/services"
)

func NewProvider(db *gorm.DB, server server.GinServer) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	routes.RegisterUserRoutes(server, userHandler)
}
