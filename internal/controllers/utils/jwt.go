package utils

import (
	"github/configs"
	"github/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(LoginData models.Users) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = LoginData.ID
	claims["username"] = LoginData.Username
	claims["email"] = LoginData.Email
	claims["phone"] = LoginData.Phone
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 3).Unix()

	process := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := process.SignedString([]byte(configs.ImportPasskey()))

	if err != nil {
		return "", err
	}

	return result, nil
}

func DecodeToken(token *jwt.Token) models.Users {
	claims := token.Claims.(jwt.MapClaims)

	var result models.Users
	result.ID = uint(claims["id"].(float64))
	result.Username = claims["username"].(string)
	result.Email = claims["email"].(string)
	result.Phone = claims["phone"].(string)

	return result
}
