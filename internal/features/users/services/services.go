package services

import (
	"github/internal/features/users"
	"github/internal/features/users/handlers"
	"github/internal/helpers"
	"github/internal/utils"
)

type UserServices struct {
	qry users.Query
}

func NewUserSer(q users.Query) users.Services {
	return &UserServices{
		qry: q,
	}
}

func (us *UserServices) Register() {
	hashedPass, err := utils.GeneratePassword(user.Password)
	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "Input Error!", nil))
	}
	user.Password = string(hashedPass)

	err = uh.srv.Register(ToRepoUsers(user))
}

func (us *UserServices) Login(user handlers.LoginRequest) {
	result, err := us.qry.Login(user.Username)

	err = utils.ComparePassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "Wrong Password!", nil))
	}

	token, err := utils.GenerateToken(result)
	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Privacy Error!", nil))
	}
}
