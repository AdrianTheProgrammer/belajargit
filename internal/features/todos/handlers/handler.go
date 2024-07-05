package handlers

import (
	"github/internal/features/todos"
	"github/internal/features/todos/repositories"
	"github/internal/helpers"
	"github/internal/utils"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TodosHand struct {
	srv todos.Services
}

func NewTodosHand(s todos.Services) todos.Handlers {
	return &TodosHand{
		srv: s,
	}
}

func (th *TodosHand) CreateTodo(c echo.Context) error {
	LoginData := utils.DecodeToken(c.Get("user").(*jwt.Token))

	var todo TodosRequest
	err := c.Bind(&todo)

	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(201, "Input Error!", nil))
	}

	err = th.srv.CreateTodo(repositories.ToTodosEntity(ToRepoTodos(todo, LoginData.ID)))

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(201, helpers.ResponseFormat(201, "Data Inserted Successfully!", nil))
}

func (th *TodosHand) ReadAllTodos(c echo.Context) error {
	LoginData := utils.DecodeToken(c.Get("user").(*jwt.Token))

	todos, err := th.srv.ReadAllTodos(LoginData.ID)

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "Activities Retrieved Successfully!", ToAllTodos(todos)))
}

func (th *TodosHand) UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	idconv, _ := strconv.Atoi(id)

	var todo TodosRequest
	err := c.Bind(&todo)

	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "Input Error!", nil))
	}

	err = th.srv.UpdateTodo(uint(idconv), TodoReqToEntity(todo))

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "Activity Updated Successfully!", nil))
}

func (th *TodosHand) DeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := th.srv.DeleteTodo(uint(id))

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "Activity Deleted Successfully!", nil))
}
