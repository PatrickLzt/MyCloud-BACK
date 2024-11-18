package main

import (
	"log"
	"os"

	"github.com/PatrickLzt/MyCloud-BACK/internal/env"
	"github.com/PatrickLzt/MyCloud-BACK/internal/store"
	"github.com/lpernett/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{
		address: env.GetString(os.Getenv("ADDRESS"), os.Getenv("PORT")),
	}

	store := store.NewPGStore(nil)

	app := &App{
		config: config,
		store:  store,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
