package users

import "github/internal/models"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func ToModelUsers(r RegisterRequest) models.Users {
	return models.Users{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
		Phone:    r.Phone,
	}
}
