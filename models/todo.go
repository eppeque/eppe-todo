package models

import "errors"

type Todo struct {
	Id     int
	Title  string
	Done   bool
	UserId int
}

func NewTodo(id int, title string, done bool, userId int) (*Todo, error) {
	if len(title) == 0 {
		return nil, errors.New("the title is empty")
	}

	return &Todo{id, title, done, userId}, nil
}
