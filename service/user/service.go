package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(registerUserRequest RegisterUserRequest) (User, error)
	Login(loginRequest LoginRequest) (User, error)
}

type service struct {
	userRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(registerUserRequest RegisterUserRequest) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerUserRequest.Password), bcrypt.MinCost)
	if err != nil {
		return User{}, err
	}

	user := User{
		Name:         registerUserRequest.Name,
		Occupation:   registerUserRequest.Occupation,
		Email:        registerUserRequest.Email,
		PasswordHash: string(passwordHash),
		Role:         "user",
	}
	return s.userRepository.Save(user)
}

func (s *service) Login(loginRequest LoginRequest) (User, error) {
	email := loginRequest.Email
	password := loginRequest.Password

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}
