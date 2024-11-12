package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiServer struct {
	address string
}

var Users = []User{}

func (api *ApiServer) getUsersHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(Users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *ApiServer) createUsersHandler(w http.ResponseWriter, r *http.Request) {

	var payload User

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	err = insertUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func insertUser(user User) error {

	// Validate user
	if user.FirstName == "" {
		return errors.New("first name is required")
	}

	if user.LastName == "" {
		return errors.New("last name is required")
	}

	// Store user validation
	for _, u := range Users {
		if u.FirstName == user.FirstName && u.LastName == user.LastName {
			return errors.New("user already exists")
		}
	}

	Users = append(Users, user)
	return nil
}
