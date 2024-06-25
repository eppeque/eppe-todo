package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type Config struct {
	Port      int              `json:"port"`
	SecretKey string           `json:"secretKey"`
	Id        int              `json:"id"`
	Users     map[string]*User `json:"users"`
}

var ServerConfig *Config

func InitConfig() *Config {
	config, err := readFromFile()

	if err == nil {
		return config
	}

	config = NewConfig()
	content, _ := json.Marshal(config)
	err = os.WriteFile("config.json", content, 0666)

	for err != nil {
		err = os.WriteFile("config.json", content, 0666)
	}

	return config
}

func NewConfig() *Config {
	return &Config{8080, generateSecretKey(), 0, make(map[string]*User)}
}

func readFromFile() (*Config, error) {
	content, err := os.ReadFile("config.json")

	if err != nil {
		return nil, err
	}

	var config = &Config{}
	err = json.Unmarshal(content, config)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) AddUser(username, email, password string) (int, error) {
	if c.Users[email] != nil {
		err := fmt.Sprintf("user with the email %s already exists", email)
		return -1, errors.New(err)
	}

	user, err := NewUser(c.Id, username, email, password)

	if err != nil {
		return -1, err
	}

	c.Users[user.Email] = user
	c.Id++
	c.saveToFile()

	return user.Id, nil
}

func (c *Config) CheckPassword(email, password string) (int, error) {
	user := c.Users[email]

	if user != nil {
		ok := user.CheckPassword(password)

		if !ok {
			return -1, errors.New("incorrect password")
		}

		return user.Id, nil
	}

	return -1, errors.New("no user found with the given email")
}

func (c *Config) UserFromId(id int) *User {
	for _, user := range c.Users {
		if user.Id == id {
			return user
		}
	}

	return nil
}

func (c *Config) saveToFile() {
	content, err := json.Marshal(c)

	if err != nil {
		return
	}

	os.WriteFile("config.json", content, 0666)
}

func generateSecretKey() string {
	result := []byte{}
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for range 22 {
		index := rand.Intn(len(chars))
		result = append(result, chars[index])
	}

	return string(result)
}
