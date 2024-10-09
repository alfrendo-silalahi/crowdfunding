package handler

import (
	"net/http"
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

	h.service.RegisterUser(registerUserRequest)

	c.JSON(http.StatusCreated, nil)
}
