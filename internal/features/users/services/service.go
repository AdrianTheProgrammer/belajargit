package services

import (
	"github/internal/features/users"
	"github/internal/utils"
)

type UsersServices struct {
	qry users.UQuery
	pu  utils.PassUtilInterface
	tu  utils.TokenUtilInterface
}

func NewUsersSer(q users.UQuery, p utils.PassUtilInterface, t utils.TokenUtilInterface) users.UServices {
	return &UsersServices{
		qry: q,
		pu:  p,
		tu:  t,
	}
}

func (us *UsersServices) Register(user users.Users) error {
	hashedPass, err := us.pu.GeneratePassword(user.Password)
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
		return users.Users{}, "", err
	}

	err = us.pu.ComparePassword([]byte(result.Password), []byte(password))
	if err != nil {
		return users.Users{}, "", err
	}

	token, err := us.tu.GenerateToken(result)
	if err != nil {
		return users.Users{}, "", err
	}

	return result, token, nil
}
