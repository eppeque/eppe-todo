package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eppeque/todo-server/infra"
	"github.com/eppeque/todo-server/models"
	"github.com/eppeque/todo-server/security"
)

type postBody struct {
	Title string `json:"title"`
}

type postResponse struct {
	Message string `json:"message"`
}

type todoDTO struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type getResponse struct {
	Todos []*todoDTO `json:"todos"`
}

func HandleTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" || r.Method == "HEAD" || r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Cannot %s /api/todos", r.Method)
		return
	}

	auth := r.Header.Get("Authorization")
	prefixLen := len("Bearer ")

	if len(auth) < prefixLen {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.NewError("You're not authorized to perform this action"))
		return
	}

	token := auth[prefixLen:]
	id, err := security.VerifyToken(token)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.NewError("You're not authorized to perform this action"))
		return
	}

	w.Header().Add("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		get(w, id)
	case "POST":
		post(w, r, id)
	case "PUT":
		put()
	case "DELETE":
		delete()
	default:
		other(w, r)
	}
}

func get(w http.ResponseWriter, userId int) {
	todos, err := models.ServerRepository.GetUserTodos(userId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewError(err.Error()))
		return
	}

	dtos := mapTodos(todos, func(todo *models.Todo) *todoDTO {
		return &todoDTO{todo.Id, todo.Title, todo.Done}
	})

	res := getResponse{dtos}
	json.NewEncoder(w).Encode(res)
}

func post(w http.ResponseWriter, r *http.Request, id int) {
	body := &postBody{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewError("The request body is invalid"))
		return
	}

	todo, err := models.ServerRepository.AddTodo(body.Title, id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewError(err.Error()))
		return
	}

	todoId, err := infra.Db.SaveTodo(id, todo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewError("An error occured while saving the todo"))
		return
	}

	todo.Id = todoId

	msg := fmt.Sprintf("Todo '%s' successfully created!", todo.Title)
	res := &postResponse{msg}
	json.NewEncoder(w).Encode(res)
}

func put() {

}

func delete() {

}

func other(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, "Cannot %s /api/todos", r.Method)
}

func mapTodos(todos []*models.Todo, convert func(todo *models.Todo) *todoDTO) (dtos []*todoDTO) {
	for _, todo := range todos {
		dtos = append(dtos, convert(todo))
	}

	return
}
