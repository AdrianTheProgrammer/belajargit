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
		return c.JSON(400, helpers.ResponseFormat(400, "Input Error!", nil))
	}

	err = uh.srv.Register(ToRepoUsers(user))

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(201, helpers.ResponseFormat(201, "Data Inserted Successfully!", nil))
}

func (uh *UsersHand) Login(c echo.Context) error {
	var user LoginRequest

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "Input Error!", nil))
	}

	result, token, err := uh.srv.Login(user.Username, user.Password)

	if err != nil {
		return c.JSON(404, helpers.ResponseFormat(404, "User Not Found!", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "Login Success!", ToLoginReponse(result, token)))
}
