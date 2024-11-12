package main

import (
	"log"
	"net/http"
	"time"
)

type App struct {
	config Config
}

type Config struct {
	address string
}

func (app *App) Mount() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", app.handleHealth)

	return mux
}

func (app *App) Run(mux *http.ServeMux) error {

	server := &http.Server{
		Addr:         app.config.address,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server is listening on %s", app.config.address)

	return server.ListenAndServe()
}
