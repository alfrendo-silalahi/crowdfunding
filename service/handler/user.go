package handler

import (
	"net/http"
	"service/helper"
	"service/user"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *handler {
	return &handler{service}
}

func (h *handler) RegisterUser(c *gin.Context) {
	var registerUserRequest user.RegisterUserRequest

	err := c.ShouldBind(&registerUserRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	newUser, err := h.service.RegisterUser(registerUserRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	data := user.MapUserToRegisterUserResponse(newUser, "token")

	response := helper.APIResponse("Account has been registered", http.StatusCreated, "Success", data)

	c.JSON(http.StatusCreated, response)
}
