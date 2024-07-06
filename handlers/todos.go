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
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type todoDTO struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type getResponse struct {
	Todos []*todoDTO `json:"todos"`
}

type putBody struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type deleteBody struct {
	Id int `json:"id"`
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
		put(w, r, id)
	case "DELETE":
		delete(w, r, id)
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

	if dtos == nil {
		dtos = []*todoDTO{}
	}

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

	res := &postResponse{todo.Id, todo.Title}
	json.NewEncoder(w).Encode(res)
}

func put(w http.ResponseWriter, r *http.Request, id int) {
	body := &putBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewError("The request body is invalid"))
		return
	}

	if !models.ServerRepository.IsTodoOwnedByUser(body.Id, id) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.NewError("You're not authorized to modify this todo"))
		return
	}

	todo, err := models.NewTodo(body.Id, body.Title, body.Done, id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewError(err.Error()))
		return
	}

	models.ServerRepository.UpdateTodo(id, todo)

	if err := infra.Db.UpdateTodo(todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewError("An error occured while updating the todo"))
		return
	}

	res := putBody{body.Id, body.Title, body.Done}
	json.NewEncoder(w).Encode(res)
}

func delete(w http.ResponseWriter, r *http.Request, id int) {
	body := &deleteBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewError("The request body is invalid"))
		return
	}

	if !models.ServerRepository.IsTodoOwnedByUser(body.Id, id) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.NewError("You're not authorized to delete this todo"))
		return
	}

	models.ServerRepository.DeleteTodo(id, body.Id)

	if err := infra.Db.DeleteTodo(body.Id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewError("An error occured while deleting the todo"))
		return
	}

	res := models.NewError("Todo deleted successfully!")
	json.NewEncoder(w).Encode(res)
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
