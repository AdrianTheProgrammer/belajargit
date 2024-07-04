package todos

import (
	"time"

	"github.com/labstack/echo"
)

type Todos struct {
	ID       uint
	Activity string
	Date     time.Time
	Status   bool
	UserID   uint
}

type Handlers interface {
	CreateTodo(echo.Context) error
	ReadAllTodos(echo.Context) error
	UpdateTodo(echo.Context) error
	DeleteTodo(echo.Context) error
}

type Services interface {
	CreateTodo(Todos) error
	ReadAllTodos(uint) ([]Todos, error)
	UpdateTodo(uint, Todos) error
	DeleteTodo(uint) error
}

type Query interface {
	CreateTodo(Todos) error
	ReadAllTodos(uint) ([]Todos, error)
	UpdateTodo(Todos) error
	DeleteTodo(uint) error
}
