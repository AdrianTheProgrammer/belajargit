package users

import "github.com/labstack/echo/v4"

type Users struct {
	ID       uint
	Username string
	Password string
	Email    string
	Phone    string
}

type UHandlers interface {
	Register(echo.Context) error
	Login(echo.Context) error
}

type UServices interface {
	Register(Users) error
	Login(string, string) (Users, string, error)
}

type UQuery interface {
	Register(Users) error
	Login(string) (Users, error)
}
