package main

import (
	"net/http"

	"github.com/eppeque/todo-server/handlers"
	"github.com/eppeque/todo-server/models"
)

func assignAPIHandlers(config *models.Config) {
	models.ServerConfig = config

	http.HandleFunc("/api/register", handlers.HandleRegister)
	http.HandleFunc("/api/login", handlers.HandleLogin)
	http.HandleFunc("/api/refresh", handlers.HandleRefresh)
	http.HandleFunc("/api/account", handlers.HandleAccount)
	http.HandleFunc("/api/todos", handlers.HandleTodos)
}
