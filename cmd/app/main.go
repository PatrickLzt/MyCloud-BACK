package main

import "log"

func main() {

	config := Config{
		address: ":8080",
	}

	app := &App{
		config: config,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
