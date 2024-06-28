package main

import (
	"fmt"
	"github/configs"
	"github/internal/controllers/todos"
	"github/internal/controllers/users"
	"github/internal/models"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := configs.ConnectDB()

	var input int
	fmt.Print("Input '1' to Migrate Database: ")
	fmt.Scanln(&input)

	if input == 1 {
		db.AutoMigrate(&models.Users{}, &models.Todos{})
	}

	um := models.NewUsersMod(db)
	tm := models.NewTodosMod(db)

	uc := users.NewUsersCon(um)
	tc := todos.NewTodosCon(tm)

	e.POST("/users/register", uc.Register)
	e.POST("/users/login", uc.Login)
	e.POST("/todos/create", tc.CreateTodo)
	e.GET("/todos/readall/:user_id", tc.ReadAllTodos)
	e.PUT("/todos/update/:id", tc.UpdateTodo)
	e.DELETE("/todos/delete/:id", tc.DeleteTodo)

	e.Start(":5000")
}
