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
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.service.RegisterUser(registerUserRequest)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := user.MapUserToRegisterUserResponse(newUser, "token")

	response := helper.APIResponse("Account has been registered", http.StatusCreated, "Success", data)

	c.JSON(http.StatusCreated, response)
}
