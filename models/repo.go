package models

import (
	"errors"
)

type Repository struct {
	users map[string]*User
	todos map[int][]*Todo
}

var ServerRepository *Repository

func NewRepository(users []*User, todos []*Todo) *Repository {
	usersMap := make(map[string]*User, len(users))
	todosMap := make(map[int][]*Todo, len(users))

	for _, user := range users {
		usersMap[user.Email] = user

		userTodos := filter(todos, func(todo *Todo) bool {
			return todo.UserId == user.Id
		})

		todosMap[user.Id] = userTodos
	}

	return &Repository{usersMap, todosMap}
}

func (r *Repository) AddUser(username, email, password string) error {
	if r.users[email] != nil {
		return errors.New("a user with the given email already exists")
	}

	user, err := NewUser(0, username, email, password)

	if err != nil {
		return err
	}

	r.users[email] = user
	return nil
}

func (r *Repository) SetIdToUser(email string, id int) error {
	if r.users[email] == nil {
		return errors.New("no user found with the given email")
	}

	r.users[email].Id = id
	r.todos[id] = []*Todo{}
	return nil
}

func (r *Repository) CheckPassword(email, password string) (int, error) {
	user := r.users[email]

	if user == nil {
		return -1, errors.New("no user found with the given email")
	}

	if user.CheckPassword(password) {
		return user.Id, nil
	}

	return -1, errors.New("the password is incorrect")
}

func (r *Repository) GetUserFromId(id int) *User {
	for _, user := range r.users {
		if user.Id == id {
			return user
		}
	}

	return nil
}

func (r *Repository) AddTodo(title string, userId int) (*Todo, error) {
	if r.todos[userId] == nil {
		return nil, errors.New("no user found with given id")
	}

	todo, err := NewTodo(0, title, false, userId)

	if err != nil {
		return nil, err
	}

	r.todos[userId] = append(r.todos[userId], todo)
	return todo, nil
}

func (r *Repository) GetUserTodos(userId int) ([]*Todo, error) {
	todos := r.todos[userId]

	if todos == nil {
		return nil, errors.New("no todos found with the given user id")
	}

	return todos, nil
}

func filter(todos []*Todo, test func(*Todo) bool) (ret []*Todo) {
	for _, todo := range todos {
		if test(todo) {
			ret = append(ret, todo)
		}
	}

	return
}
