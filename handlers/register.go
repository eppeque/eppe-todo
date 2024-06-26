package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eppeque/todo-server/infra"
	"github.com/eppeque/todo-server/models"
	"github.com/eppeque/todo-server/security"
)

type registerBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerResponse struct {
	Token string `json:"token"`
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Cannot %s /api/register", r.Method)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	body := &registerBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewError("The request body is invalid"))
		return
	}

	if err := models.ServerRepository.AddUser(body.Username, body.Email, body.Password); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewError(err.Error()))
		return
	}

	user, _ := models.NewUser(0, body.Username, body.Email, body.Password)
	id, err := infra.Db.SaveUser(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewError("Something went wrong while saving the user"))
		return
	}

	models.ServerRepository.SetIdToUser(user.Email, id)
	token := security.CreateToken(id)
	res := &registerResponse{token}
	json.NewEncoder(w).Encode(res)
}
