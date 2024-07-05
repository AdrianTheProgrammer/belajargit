package main

import (
	"fmt"
	"github/configs"
	t_hnd "github/internal/features/todos/handlers"
	t_rep "github/internal/features/todos/repositories"
	t_srv "github/internal/features/todos/services"
	u_hnd "github/internal/features/users/handlers"
	u_rep "github/internal/features/users/repositories"
	u_srv "github/internal/features/users/services"
	"github/internal/utils"

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
		db.AutoMigrate(&u_rep.Users{}, &t_rep.Todos{})
	}
	pu := utils.NewPassUtil()
	tu := utils.NewTokenUtil()
	uq := u_rep.NewUsersQue(db)
	us := u_srv.NewUsersSer(uq, pu, tu)
	uh := u_hnd.NewUsersHand(us)

	tq := t_rep.NewTodosQue(db)
	ts := t_srv.NewTodosSer(tq)
	th := t_hnd.NewTodosHand(ts, tu)

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
