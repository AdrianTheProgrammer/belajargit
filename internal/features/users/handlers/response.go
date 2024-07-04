package handlers

import "github/internal/features/users"

type LoginResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Token    string `json:"token"`
}

func ToLoginReponse(user users.Users, token string) LoginResponse {
	return LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
		Token:    token,
	}
}
