package todos

import (
	"github/internal/helpers"
	"github/internal/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodosCon struct {
	model *models.TodosMod
}

func NewTodosCon(m *models.TodosMod) *TodosCon {
	return &TodosCon{
		model: m,
	}
}

func (uc *TodosCon) CreateTodo(c echo.Context) error {
	var user models.Todos
	err := c.Bind(&user)

	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(201, "Input Error!", nil))
	}

	err = uc.model.CreateTodo(user)

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(201, helpers.ResponseFormat(201, "Data Inserted Successfully!", nil))
}

func (uc *TodosCon) ReadAllTodos(c echo.Context) error {
	userID := c.Param("user_id")
	var todos []models.Todos
	err := c.Bind(&todos)

	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "Input Error!", nil))
	}

	todos, err = uc.model.ReadAllTodos(userID)

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "Activities Retrieved Successfully!", ToAllTodos(todos)))
}

func (uc *TodosCon) UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	idconv, _ := strconv.Atoi(id)

	var todo models.Todos
	todo.ID = uint(idconv)
	err := c.Bind(&todo)

	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "Input Error!", nil))
	}

	err = uc.model.UpdateTodo(todo)

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "Activity Updated Successfully!", nil))
}

func (uc *TodosCon) DeleteTodo(c echo.Context) error {
	id := c.Param("id")

	err := uc.model.DeleteTodo(id)

	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "Activity Deleted Successfully!", nil))
}
