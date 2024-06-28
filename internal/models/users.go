package models

import (
	"gorm.io/gorm"
)

type UsersMod struct {
	db *gorm.DB
}

func NewUsersMod(connection *gorm.DB) *UsersMod {
	return &UsersMod{
		db: connection,
	}
}

type Users struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (um *UsersMod) Register(user Users) error {
	err := um.db.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (um *UsersMod) Login(user Users) (Users, error) {
	var result Users
	err := um.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}
