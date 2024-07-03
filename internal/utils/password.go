package utils

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(plainPass string) ([]byte, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(plainPass), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ComparePassword(currentPass, inputPass []byte) error {
	return bcrypt.CompareHashAndPassword(currentPass, inputPass)
}
