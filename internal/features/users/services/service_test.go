package services_test

import (
	"github/internal/features/users"
	"github/internal/features/users/services"
	"github/mocks"
	"testing"
)

func TestRegister(t *testing.T) {
	qry := mocks.NewQuery(t)
	srv := services.NewUsersSer(qry)

	input := users.Users{
		ID:       1,
		Username: "admin",
		Password: "admin",
		Email:    "admin@gmail.com",
		Phone:    "089123456879",
	}

	srv.Register(input)
}
