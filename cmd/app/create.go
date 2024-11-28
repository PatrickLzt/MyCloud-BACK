package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/PatrickLzt/MyCloud-BACK/internal/store"
)

func (app *App) handleCreate(w http.ResponseWriter, r *http.Request) {

	var user store.User

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading the request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Error parsing the request body", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Println("POST /users", user)

}
