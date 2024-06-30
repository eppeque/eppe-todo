package infra

import (
	"errors"
	"math/rand"
	"os"
	"strings"
)

var SecretKey string

func initSecret() (string, error) {
	secret, err := readFromFile()

	if err == nil {
		return secret, nil
	}

	secret = newSecret()
	err = os.WriteFile("secret.txt", []byte(secret), 0666)

	if err != nil {
		return "", err
	}

	return secret, nil
}

func newSecret() string {
	result := make([]byte, 22)
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for i := range 22 {
		index := rand.Intn(len(chars))
		result[i] = chars[index]
	}

	return string(result)
}

func readFromFile() (string, error) {
	content, err := os.ReadFile("secret.txt")

	if err != nil {
		return "", err
	}

	secret := string(content)
	secret = strings.TrimSpace(secret)

	if len(secret) != 22 {
		return "", errors.New("invalid secret key")
	}

	return secret, nil
}
