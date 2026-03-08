package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"service/helper"
	"service/user"
)

type handler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *handler {
	return &handler{service}
}

func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var registerUserRequest user.RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&registerUserRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.RegisterUser(registerUserRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("User successfully registered."))
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest user.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.service.Login(loginRequest)
	if err != nil {
		log.Printf("Login error: %v", err)

		resMessage := helper.APIResponse(
			"Invalid email or password.",
			http.StatusUnauthorized,
			"error",
			nil,
		)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resMessage)
		return
	}

	w.Write([]byte("User successfully logged in."))
}
