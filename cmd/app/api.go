package main

import (
	"log"
	"net/http"
	"time"

	"github.com/PatrickLzt/MyCloud-BACK/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	config Config
	store  store.Storage
}

type Config struct {
	address string
}

func (app *App) Mount() http.Handler {

	router := chi.NewRouter() // Create a Chi router instance

	router.Use(middleware.RequestID) // Use the Chi request ID middleware
	router.Use(middleware.RealIP)    // Use the Chi real IP middleware
	router.Use(middleware.Recoverer) // Use the Chi recoverer middleware
	router.Use(middleware.Logger)    // Use the Chi logger middleware

	router.Use(middleware.Timeout(30 * time.Second)) // Use the Chi timeout middleware

	router.Get("/health", app.handleHealth)

	return router
}

func (app *App) Run(mux http.Handler) error {

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
