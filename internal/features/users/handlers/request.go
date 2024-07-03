package handlers

import "github/internal/features/users"

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

func ToRepoUsers(r RegisterRequest) users.Users {
	return users.Users{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
		Phone:    r.Phone,
	}
}
