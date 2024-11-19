package users

import "github.com/labstack/echo/v4"

type Users struct {
	ID       uint
	Username string
	Password string
	Email    string
	Phone    string
}

type Handlers interface {
	Register(echo.Context) error
	Login(echo.Context) error
}

type Services interface {
	Register(Users) (int, error)
	Login(string, string) (Users, string, int, error)
}

type Query interface {
	Register(Users) error
	Login(string) (Users, error)
}

type RegisterValidate struct {
	Username string `validate:"required"`
	Password string `validate:"required,min=6,alphanum"`
	Email    string `validate:"required"`
	Phone    string `validate:"required"`
}

type LoginValidate struct {
	Username string `validate:"required"`
	Password string `validate:"required,min=6,alphanum"`
}
