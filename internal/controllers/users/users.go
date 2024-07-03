package users

import (
	"github/internal/helpers"
	"github/internal/models"
	"github/internal/utils"

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

	hashedPass, err := utils.GeneratePassword(user.Password)
	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "Input Error!", nil))
	}
	user.Password = string(hashedPass)

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

	result, err := uc.model.Login(user.Username)
	if err != nil {
		return c.JSON(404, helpers.ResponseFormat(404, "User Not Found!", nil))
	}

	err = utils.ComparePassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "Wrong Password!", nil))
	}

	token, err := utils.GenerateToken(result)
	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Privacy Error!", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "Login Success!", ToLoginReponse(result, token)))
}
