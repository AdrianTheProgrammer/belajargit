package repositories

import (
	"github/internal/features/users"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (u *Users) ToUsersEntity() users.Users {
	return users.Users{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Phone:    u.Phone,
	}
}

func ToUsersData(input users.Users) Users {
	return Users{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	}
}
