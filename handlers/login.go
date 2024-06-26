package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eppeque/todo-server/models"
	"github.com/eppeque/todo-server/security"
)

type loginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Cannot %s /api/login", r.Method)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	body := &loginBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewError("The request body is invalid"))
		return
	}

	id, err := models.ServerRepository.CheckPassword(body.Email, body.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewError("Incorrect email and/or password"))
		return
	}

	token := security.CreateToken(id)
	res := &loginResponse{token}
	json.NewEncoder(w).Encode(res)
}
