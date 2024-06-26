package models

import "errors"

type Repository struct {
	users map[string]*User
}

var ServerRepository *Repository

func NewRepository(users []*User) *Repository {
	usersMap := make(map[string]*User, len(users))

	for _, user := range users {
		usersMap[user.Email] = user
	}

	return &Repository{usersMap}
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
