package main

import (
	"fmt"
	"github/configs"
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
		db.AutoMigrate(&models.Users{})
	}

	um := models.NewUsersMod(db)
	uc := users.NewUsersCon(um)

	e.POST("/register", uc.Register)
	e.POST("/login", uc.Login)

	e.Start(":5000")
}
