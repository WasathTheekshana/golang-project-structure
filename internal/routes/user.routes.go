package routes

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/WasathTheekshana/golang-project-structure/cmd/server"
	"github.com/WasathTheekshana/golang-project-structure/interfaces"
	"github.com/WasathTheekshana/golang-project-structure/internal/handler"
)

func RegisterUserRoutes(server server.GinServer, userHandler *handler.UserHandler) {
	server.RegisterGroupRoute("/user", []interfaces.IRouteDefinition{
		{
			Method:  "POST",
			Path:    "/register",
			Handler: userHandler.CreateUser,
		},
	}, func(c *gin.Context) {
		log.Info("Request on %s", c.Request.URL.Path)
	})
}
