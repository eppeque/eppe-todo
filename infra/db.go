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

	_, err = database.Exec("create table if not exists users (username text not null, email text not null unique, password text not null);")

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

func (d *ServerDatabase) GetAllUsers() ([]*models.User, error) {
	var users []*models.User

	rows, err := d.db.Query("select rowid, username, email, password from users;")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
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
