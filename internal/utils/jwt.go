package utils

import (
	"github/configs"
	"github/internal/features/users"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(LoginData users.Users) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = LoginData.ID
	claims["username"] = LoginData.Username
	claims["password"] = LoginData.Password
	claims["email"] = LoginData.Email
	claims["phone"] = LoginData.Phone
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	process := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := process.SignedString([]byte(configs.ImportPasskey()))

	if err != nil {
		return "", err
	}

	return result, nil
}

func DecodeToken(token *jwt.Token) users.Users {
	claims := token.Claims.(jwt.MapClaims)

	var result users.Users
	result.ID = uint(claims["id"].(float64))
	result.Username = claims["username"].(string)
	result.Password = claims["password"].(string)
	result.Email = claims["email"].(string)
	result.Phone = claims["phone"].(string)

	return result
}
