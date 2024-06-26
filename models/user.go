package models

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Username string
	Email    string
	Password string
}

func NewUser(id int, username, email, password string) (*User, error) {
	usernameRx := regexp.MustCompile(`^[a-zA-Z0-9_]{5,}$`)
	emailRx := regexp.MustCompile(`^[a-z0-9]+(\.[a-z0-9]+)*@[a-z0-9]+(\.[a-z0-9]+)*\.[a-z]+$`)
	passwordRx := regexp.MustCompile(`^.{8,}$`)

	if !usernameRx.Match([]byte(username)) {
		return nil, errors.New("username has an invalid format")
	}

	if !emailRx.Match([]byte(email)) {
		return nil, errors.New("email has an invalid format")
	}

	if !passwordRx.Match([]byte(password)) {
		return nil, errors.New("password is not long enough")
	}

	hashed, err := hashPassword(password)

	if err != nil {
		return nil, err
	}

	return &User{id, username, email, hashed}, nil
}

func (u *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}
