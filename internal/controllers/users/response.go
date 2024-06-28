package users

import "github/internal/models"

type LoginResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func ToLoginReponse(user models.Users) LoginResponse {
	return LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
	}
}
