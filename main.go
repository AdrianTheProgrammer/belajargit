package main

import (
	"fmt"
	"github/configs"
	"github/internal/controllers/todos"
	"github/internal/controllers/users"
	"github/internal/models"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	db := configs.ConnectDB()
	passkey := configs.ImportPasskey()

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

	// USERS
	e.POST("/users/register", uc.Register)
	e.POST("/users/login", uc.Login)

	// TODOS
	t := e.Group("/todos")
	t.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(passkey),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))
	t.POST("/create", tc.CreateTodo)
	t.GET("/read_all", tc.ReadAllTodos)
	t.PUT("/update/:id", tc.UpdateTodo)
	t.DELETE("/delete/:id", tc.DeleteTodo)

	// MIDDLEWARE
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// START
	e.Logger.Error(e.Start(":5000"))
}
