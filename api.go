package main

import (
	"log"
	"net/http"

	"github.com/eppeque/todo-server/handlers"
	"github.com/eppeque/todo-server/infra"
	"github.com/eppeque/todo-server/models"
)

func assignAPIHandlers() {
	users, err := infra.Db.GetAllUsers()

	if err != nil {
		log.Fatalln(err)
	}

	models.ServerRepository = models.NewRepository(users)

	http.HandleFunc("/api/register", handlers.HandleRegister)
	http.HandleFunc("/api/login", handlers.HandleLogin)
	http.HandleFunc("/api/refresh", handlers.HandleRefresh)
	http.HandleFunc("/api/account", handlers.HandleAccount)
	http.HandleFunc("/api/todos", handlers.HandleTodos)
}
