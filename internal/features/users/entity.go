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
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type Services interface {
	Register(user Users) error
	Login(username string, password string) (Users, string, error)
}

type Query interface {
	Register(user Users) error
	Login(username string) (Users, error)
}
