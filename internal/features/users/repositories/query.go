package repositories

import (
	"github/internal/features/users"

	"gorm.io/gorm"
)

type UsersQue struct {
	db *gorm.DB
}

func NewUsersQue(connection *gorm.DB) *UsersQue {
	return &UsersQue{
		db: connection,
	}
}

func (uq *UsersQue) Register(user users.Users) error {
	cnv := ToUsersData(user)
	err := uq.db.Create(&cnv).Error

	if err != nil {
		return err
	}

	return nil
}

func (uq *UsersQue) Login(username string) (Users, error) {
	var result Users
	err := uq.db.Where("username = ?", username).First(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}
