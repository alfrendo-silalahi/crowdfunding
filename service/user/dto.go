package user

type RegisterUserRequest struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type registerUserResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func MapUserToRegisterUserResponse(user User, token string) registerUserResponse {
	return registerUserResponse{
		Id:         user.Id,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
}
