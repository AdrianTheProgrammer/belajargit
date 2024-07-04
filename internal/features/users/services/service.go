package services

import (
	"github/internal/features/users"
	"github/internal/utils"
)

type UsersServices struct {
	qry users.Query
}

func NewUsersSer(q users.Query) users.Services {
	return &UsersServices{
		qry: q,
	}
}

func (us *UsersServices) Register(user users.Users) error {
	hashedPass, err := utils.GeneratePassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)

	err = us.qry.Register(user)

	return err
}

func (us *UsersServices) Login(username, password string) (users.Users, string, error) {
	result, err := us.qry.Login(username)
	if err != nil {
		return users.Users{}, "", nil
	}

	err = utils.ComparePassword([]byte(result.Password), []byte(password))
	if err != nil {
		return users.Users{}, "", nil
	}

	token, err := utils.GenerateToken(result)
	if err != nil {
		return users.Users{}, "", nil
	}

	return result, token, nil
}
