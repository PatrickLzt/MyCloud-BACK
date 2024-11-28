package main

import (
	"log"
	"os"

	"github.com/PatrickLzt/MyCloud-BACK/internal/db"
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
		db: dbConfig{
			address:            env.GetString(os.Getenv("DB_ADDRESS"), "postgres://postgres:21622292a@localhost:5432/db_test?sslmode=disable"),
			maxOpenConnections: env.GetInt(os.Getenv("DB_MAX_OPEN_CONNECTIONS"), 10),
			maxIdleConnections: env.GetInt(os.Getenv("DB_MAX_IDLE_CONNECTIONS"), 10),
			maxIdleTime:        env.GetString(os.Getenv("DB_MAX_IDLE_TIME"), "15m"),
		},
	}

	db, err := db.New(config.db.address, config.db.maxOpenConnections, config.db.maxIdleConnections, config.db.maxIdleTime)

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	defer db.Close()
	log.Printf("Connected to the database")

	store := store.NewPGStore(db)

	app := &App{
		config: config,
		store:  store,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
