package handlers

import (
	"github/internal/features/users"
	"github/internal/helpers"

	"github.com/labstack/echo/v4"
)

type UsersHand struct {
	srv users.Services
}

func NewUsersHand(s users.Services) users.Handlers {
	return &UsersHand{
		srv: s,
	}
}

func (uh *UsersHand) Register(c echo.Context) error {
	var user RegisterRequest

	err := c.Bind(&user)
	if err != nil {
		c.Logger().Error("register parse error:", err.Error())
		return c.JSON(400, helpers.ResponseFormat(400, "Input Error!", nil))
	}

	errCode, errMsg := uh.srv.Register(ToRepoUsers(user))
	if errMsg != nil {
		return c.JSON(errCode, helpers.ResponseFormat(errCode, errMsg.Error(), nil))
	}

	return c.JSON(201, helpers.ResponseFormat(201, "Data Inserted Successfully!", nil))
}

func (uh *UsersHand) Login(c echo.Context) error {
	var user LoginRequest

	err := c.Bind(&user)
	if err != nil {
		c.Logger().Error("login parse error:", err.Error())
		return c.JSON(400, helpers.ResponseFormat(400, "Input Error!", nil))
	}

	result, token, errCode, errMsg := uh.srv.Login(user.Username, user.Password)
	if errMsg != nil {
		return c.JSON(errCode, helpers.ResponseFormat(errCode, errMsg.Error(), nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "Login Success!", ToLoginReponse(result, token)))
}
