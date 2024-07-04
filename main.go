package main

import (
	"fmt"
	"github/configs"
	todoshand "github/internal/features/todos/handlers"
	todosrepo "github/internal/features/todos/repositories"
	todosserv "github/internal/features/todos/services"
	usershand "github/internal/features/users/handlers"
	usersrepo "github/internal/features/users/repositories"
	usersserv "github/internal/features/users/services"

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
		db.AutoMigrate(&usersrepo.Users{}, &todosrepo.Todos{})
	}

	uq := usersrepo.NewUsersQue(db)
	us := usersserv.NewUsersSer(uq)
	uh := usershand.NewUsersHand(us)

	tq := todosrepo.NewTodosQue(db)
	ts := todosserv.NewTodosSer(tq)
	th := todoshand.NewTodosHand(ts)

	// USERS
	e.POST("/users/register", uh.Register)
	e.POST("/users/login", uh.Login)

	// TODOS
	t := e.Group("/todos")
	t.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(passkey),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))
	t.POST("/create", th.CreateTodo)
	t.GET("/read_all", th.ReadAllTodos)
	t.PUT("/update/:id", th.UpdateTodo)
	t.DELETE("/delete/:id", th.DeleteTodo)

	// MIDDLEWARE
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// START
	e.Logger.Error(e.Start(":5000"))
}
