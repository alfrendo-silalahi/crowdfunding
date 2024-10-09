package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(registerUserRequest RegisterUserRequest) (User, error)
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
