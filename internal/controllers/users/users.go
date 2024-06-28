package users

import (
	"github/internal/helpers"
	"github/internal/models"

	"github.com/labstack/echo/v4"
)

type UsersCon struct {
	model *models.UsersMod
}

func NewUsersCon(m *models.UsersMod) *UsersCon {
	return &UsersCon{
		model: m,
	}
}

func (uc *UsersCon) Register(c echo.Context) error {
	var user models.Users
	err := c.Bind(&user)

	if err != nil {
		return c.JSON(400, "Input Error!")
	}

	err = uc.model.Register(user)

	if err != nil {
		return c.JSON(500, "Server Error!")
	}

	return c.JSON(201, "Data Inserted Successfully!")
}

func (uc *UsersCon) Login(c echo.Context) error {
	var user models.Users
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(400, "Input Error!")
	}
	result, err := uc.model.Login(user)

	if err != nil {
		return c.JSON(500, "Server Error!")
	}

	return c.JSON(200, helpers.ResponseFormat(200, "success login", ToLoginReponse(result)))
}
