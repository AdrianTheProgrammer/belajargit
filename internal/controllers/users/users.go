package users

import (
	"github/internal/controllers/utils"
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
	var user RegisterRequest

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "Input Error!", nil))
	}

	err = uc.model.Register(ToModelUsers(user))
	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(201, helpers.ResponseFormat(201, "Data Inserted Successfully!", nil))
}

func (uc *UsersCon) Login(c echo.Context) error {
	var user LoginRequest

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "Input Error!", nil))
	}

	result, err := uc.model.Login(user.Username, user.Password)
	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error!", nil))
	}

	token, err := utils.GenerateToken(result)
	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Privacy Error!", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "Login Success!", ToLoginReponse(result, token)))
}
