package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/WasathTheekshana/golang-project-structure/interfaces"
	"github.com/WasathTheekshana/golang-project-structure/internal/services"
)

type UserHandler struct {
	userService services.UserService
	validate    *validator.Validate
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		validate:    validator.New(),
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userRequest interfaces.IUserRegistrationRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadGateway, interfaces.IInfoMessage{
			Message: err.Error(),
			Status:  interfaces.StatusError,
			Code:    http.StatusBadGateway,
		})
		return
	}

	if err := h.validate.Struct(userRequest); err != nil {
		c.JSON(http.StatusBadRequest, interfaces.IInfoMessage{
			Message: err.Error(),
			Status:  interfaces.StatusError,
			Code:    http.StatusBadGateway,
		})
		return
	}

	userData, err := h.userService.CreateUserAccount(&userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfaces.IInfoMessage{
			Message: err.Error(),
			Status:  interfaces.StatusError,
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, interfaces.IUserResponse{
		Info: interfaces.IInfoMessage{
			Status: interfaces.StatusSuccess,
			Code:   http.StatusOK,
		},
		Data: *userData,
	})
}
