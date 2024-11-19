package services

import (
	"errors"
	"fmt"
	"github/internal/features/users"
	"github/internal/utils"
	"log"

	"github.com/go-playground/validator/v10"
)

type UsersServices struct {
	qry      users.Query
	validate *validator.Validate
}

func NewUsersSer(q users.Query) users.Services {
	return &UsersServices{
		qry:      q,
		validate: validator.New(),
	}
}

func (us *UsersServices) Register(user users.Users) (int, error) {
	err := us.validate.Struct(&users.RegisterValidate{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
	})
	if err != nil {
		log.Println("Invalid Input. Error:", err)
		return 400, errors.New("invalid input")
	}

	hashedPass, err := utils.GeneratePassword(user.Password)
	if err != nil {
		log.Println("Error generating password in register process:", err)
		return 400, errors.New("register failed, invalid data")
	}
	user.Password = string(hashedPass)

	err = us.qry.Register(user)
	if err != nil {
		log.Println("SQL Error in register process:", err)
		return 500, errors.New("server error when processing data")
	}

	return 201, err
}

func (us *UsersServices) Login(username, password string) (users.Users, string, int, error) {
	err := us.validate.Struct(&users.LoginValidate{Username: username, Password: password})
	if err != nil {
		log.Println("Invalid Input. Error:", err)
		return users.Users{}, "", 400, errors.New("invalid input")
	}

	result, err := us.qry.Login(username)
	if err != nil {
		log.Printf("User '%v' not found in database. Error: %v\n", username, err)
		return users.Users{}, "", 404, fmt.Errorf("user '%v' not found", username)
	}

	err = utils.ComparePassword([]byte(result.Password), []byte(password))
	if err != nil {
		log.Printf("Wrong password for user '%v'. Error: %v\n", username, err)
		return users.Users{}, "", 401, fmt.Errorf("wrong password for user '%v'", username)
	}

	token, err := utils.GenerateToken(result)
	if err != nil {
		log.Printf("Error generating JWT for user '%v': %v\n", username, err)
		return users.Users{}, "", 500, fmt.Errorf("error generating jwt for user '%v'", username)
	}

	return result, token, 200, nil
}
