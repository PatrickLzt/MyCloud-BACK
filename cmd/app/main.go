package main

import (
	"log"
	"os"

	"github.com/PatrickLzt/MyCloud-BACK/internal/env"
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

	app := &App{
		config: config,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
