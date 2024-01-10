package main

import (
	"backend/cmd/config"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	env := config.LoadEnv()

	// set application config
	var app application

	// read from command line

	// connect to the database
	app.Domain = "example.com"

	log.Println("Starting application on port", env.ApiPort)

	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
