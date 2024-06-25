package handlers

import (
	"fmt"
	"net/http"
)

func HandleTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Todos!")
}
