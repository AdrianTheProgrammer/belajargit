package services_test

import (
	"github/internal/features/users"
	"github/internal/features/users/handlers"
	"github/internal/features/users/services"
	"github/mocks"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestRegister(t *testing.T) {
	qry := mocks.NewQuery(t)
	pu := mocks.NewPassUtilInterface(t)
	tu := mocks.NewTokenUtilInterface(t)
	srv := services.NewUsersSer(qry, pu, tu)

	input := users.Users{
		ID:       1,
		Username: "admin",
		Password: "admin",
		Email:    "admin@gmail.com",
		Phone:    "089123456879",
	}

	t.Run("Password Error", func(t *testing.T) {
		pu.On("GeneratePassword", input.Password).Return(nil, bcrypt.ErrPasswordTooLong).Once()
		err := srv.Register(input)

		pu.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Query Error", func(t *testing.T) {
		inputQry := users.Users{
			ID:       1,
			Username: "admin",
			Password: "hashed_admin",
			Email:    "admin@gmail.com",
			Phone:    "089123456879",
		}
		pu.On("GeneratePassword", input.Password).Return([]byte("hashed_admin"), nil).Once()
		qry.On("Register", inputQry).Return(gorm.ErrInvalidData).Once()
		err := srv.Register(input)

		pu.AssertExpectations(t)
		qry.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Register Success", func(t *testing.T) {
		inputQry := users.Users{
			ID:       1,
			Username: "admin",
			Password: "hashed_admin",
			Email:    "admin@gmail.com",
			Phone:    "089123456879",
		}
		pu.On("GeneratePassword", input.Password).Return([]byte("hashed_admin"), nil).Once()
		qry.On("Register", inputQry).Return(nil).Once()
		err := srv.Register(input)

		pu.AssertExpectations(t)
		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestLogin(t *testing.T) {
	qry := mocks.NewQuery(t)
	pu := mocks.NewPassUtilInterface(t)
	tu := mocks.NewTokenUtilInterface(t)
	srv := services.NewUsersSer(qry, pu, tu)

	input := handlers.LoginRequest{
		Username: "admin",
		Password: "hashed_admin",
	}

	loginData := users.Users{
		ID:       1,
		Username: "admin",
		Password: "hashed_admin",
		Email:    "admin@gmail.com",
		Phone:    "089123456879",
	}

	t.Run("Query Error", func(t *testing.T) {
		qry.On("Login", input.Username).Return(users.Users{}, gorm.ErrInvalidData).Once()
		_, _, err := srv.Login(input.Username, input.Password)

		qry.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Password Error", func(t *testing.T) {
		qry.On("Login", input.Username).Return(loginData, nil).Once()
		pu.On("ComparePassword", []byte(loginData.Password), []byte("hashed_admin")).Return(bcrypt.ErrMismatchedHashAndPassword).Once()
		_, _, err := srv.Login(input.Username, input.Password)

		qry.AssertExpectations(t)
		pu.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Token Error", func(t *testing.T) {
		qry.On("Login", input.Username).Return(loginData, nil).Once()
		pu.On("ComparePassword", []byte(loginData.Password), []byte("hashed_admin")).Return(nil).Once()
		tu.On("GenerateToken", loginData).Return("", jwt.ErrTokenInvalidClaims).Once()
		_, _, err := srv.Login(input.Username, input.Password)

		qry.AssertExpectations(t)
		pu.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Login Success", func(t *testing.T) {
		token := "hasil_token"

		qry.On("Login", input.Username).Return(loginData, nil).Once()
		pu.On("ComparePassword", []byte(loginData.Password), []byte("hashed_admin")).Return(nil).Once()
		tu.On("GenerateToken", loginData).Return(token, nil).Once()

		_, _, err := srv.Login(input.Username, input.Password)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}
