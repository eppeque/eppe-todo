package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eppeque/todo-server/models"
	"github.com/eppeque/todo-server/security"
)

type refreshBody struct {
	Token string `json:"token"`
}

type refreshResponse = refreshBody

func HandleRefresh(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Cannot %s /api/refresh", r.Method)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	body := &refreshBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewError("The request body is invalid"))
		return
	}

	id, err := security.VerifyToken(body.Token)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewError("The token is invalid or expired"))
		return
	}

	newToken := security.CreateToken(id)
	res := &refreshResponse{newToken}
	json.NewEncoder(w).Encode(res)

}
