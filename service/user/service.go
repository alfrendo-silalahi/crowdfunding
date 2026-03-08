package user

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(registerUserRequest RegisterUserRequest) error
	Login(loginRequest LoginRequest) (User, error)
}

type service struct {
	userRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(registerUserRequest RegisterUserRequest) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerUserRequest.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user := User{
		Username:   registerUserRequest.Username,
		Occupation: registerUserRequest.Occupation,
		Email:      registerUserRequest.Email,
		Password:   string(passwordHash),
		Role:       "user",
	}

	err = s.userRepository.Save(user)
	if err != nil {
		log.Printf("Error saving user: %v", err)
		return errors.New("Failed to register new user.")
	}

	return nil
}

func (s *service) Login(loginRequest LoginRequest) (User, error) {
	email := loginRequest.Email
	password := loginRequest.Password

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}
