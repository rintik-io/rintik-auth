package common

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func GenerateRandPassword(passwordLength int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, passwordLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateBcryptPassword(password string, cost int) (string, error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(passwordHashed), nil
}
