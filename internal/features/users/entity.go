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
	Register(Users) error
	Login(string, string) (Users, string, error)
}

type Query interface {
	Register(Users) error
	Login(string) (Users, error)
}
