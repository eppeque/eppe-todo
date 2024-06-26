package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eppeque/todo-server/models"
	"github.com/eppeque/todo-server/security"
)

type accountResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func HandleAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Cannot %s /api/account", r.Method)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	auth := r.Header.Get("Authorization")

	if len(auth) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.NewError("You're not authorized to get this resource"))
		return
	}

	token := auth[len("Bearer "):]
	id, err := security.VerifyToken(token)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.NewError("You're not authorized to get this resource"))
		return
	}

	user := models.ServerRepository.GetUserFromId(id)

	if user == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewError("No account found"))
		return
	}

	res := accountResponse{user.Username, user.Email}
	json.NewEncoder(w).Encode(res)
}
