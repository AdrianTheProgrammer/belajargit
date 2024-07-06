package todos

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Todos struct {
	ID       uint
	Activity string
	Date     time.Time
	Status   bool
	UserID   uint
}

type THandlers interface {
	CreateTodo(echo.Context) error
	ReadAllTodos(echo.Context) error
	UpdateTodo(echo.Context) error
	DeleteTodo(echo.Context) error
}

type TServices interface {
	CreateTodo(Todos) error
	ReadAllTodos(uint) ([]Todos, error)
	UpdateTodo(uint, Todos) error
	DeleteTodo(uint) error
}

type TQuery interface {
	CreateTodo(Todos) error
	ReadAllTodos(uint) ([]Todos, error)
	UpdateTodo(uint, Todos) error
	DeleteTodo(uint) error
}
