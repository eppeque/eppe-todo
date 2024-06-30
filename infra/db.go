package infra

import (
	"database/sql"
	"os"

	"github.com/eppeque/todo-server/models"
	_ "github.com/mattn/go-sqlite3"
)

type ServerDatabase struct {
	db *sql.DB
}

var Db *ServerDatabase

func createDatabase() (*ServerDatabase, error) {
	if err := createFileIfNotExists(); err != nil {
		return nil, err
	}

	database, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		return nil, err
	}

	database.SetMaxOpenConns(1)

	_, err = database.Exec("create table if not exists users (username text not null, email text not null unique, password text not null); create table if not exists todos (title text not null, done integer not null, user integer not null, foreign key(user) references user(rowid));")

	if err != nil {
		return nil, err
	}

	return &ServerDatabase{database}, nil
}

func (d *ServerDatabase) Close() {
	d.db.Close()
}

func (d *ServerDatabase) SaveUser(user *models.User) (int, error) {
	res, err := d.db.Exec("insert into users values (?, ?, ?);", user.Username, user.Email, user.Password)

	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func (d *ServerDatabase) GetAllUsers() (users []*models.User, err error) {
	rows, err := d.db.Query("select rowid, username, email, password from users;")

	if err != nil {
		users = nil
		return
	}

	for rows.Next() {
		var user models.User

		if err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password); err != nil {
			users = nil
			return
		}

		users = append(users, &user)
	}

	return
}

func (d *ServerDatabase) SaveTodo(userId int, todo *models.Todo) (int, error) {
	res, err := d.db.Exec("insert into todos values (?, ?, ?);", todo.Title, todo.Done, userId)

	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func (d *ServerDatabase) GetAllTodos() (todos []*models.Todo, err error) {
	rows, err := d.db.Query("select rowid, title, done, user from todos;")

	if err != nil {
		todos = nil
		return
	}

	for rows.Next() {
		var todo models.Todo

		if err = rows.Scan(&todo.Id, &todo.Title, &todo.Done, &todo.UserId); err != nil {
			todos = nil
			return
		}

		todos = append(todos, &todo)
	}

	return
}

func (d *ServerDatabase) UpdateTodo(todo *models.Todo) error {
	_, err := d.db.Exec("update todos set title = ?, done = ? where rowid = ?", todo.Title, todo.Done, todo.Id)
	return err
}

func CloseDatabase() {
	Db.Close()
}

func createFileIfNotExists() error {
	_, err := os.Open("data.db")

	if err == nil {
		return nil
	}

	_, err = os.Create("data.db")
	return err
}
