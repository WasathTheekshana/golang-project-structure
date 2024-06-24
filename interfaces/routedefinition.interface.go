package interfaces

import "github.com/gin-gonic/gin"

type IRouteDefinition struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}
