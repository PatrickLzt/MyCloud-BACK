package main

import (
	"net/http"
)

func main() {

	api := &ApiServer{address: ":8080"}

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    api.address,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
