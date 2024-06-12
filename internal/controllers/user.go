package controllers

import (
	"errors"
	"fmt"
	"github/internal/models"
)

type UserController struct {
	model *models.UserModel
}

func NewUserController(m *models.UserModel) *UserController {
	return &UserController{
		model: m,
	}
}

func (uc *UserController) Login() (models.User, error) {
	var email, password string
	fmt.Println("=== LOGIN ===")
	fmt.Print("Masukkan Email: ")
	fmt.Scanln(&email)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&password)
	fmt.Println()
	result, err := uc.model.Login(email, password)
	if err != nil {
		return models.User{}, errors.New("terjadi masalah ketika login")
	}
	return result, nil
}

func (uc *UserController) Register() (models.User, error) {
	var newData models.User
	fmt.Println("=== REGISTER ===")
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&newData.Name)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&newData.Password)
	fmt.Print("Masukkan Email: ")
	fmt.Scanln(&newData.Email)
	fmt.Print("Masukkan HP: ")
	fmt.Scanln(&newData.Phone)
	result, err := uc.model.Register(newData)
	if err != nil && !result {
		return models.User{}, errors.New("terjadi masalah ketika registrasi")
	}
	return newData, nil
}
