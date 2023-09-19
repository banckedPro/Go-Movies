package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// Set application config
	var app application

	// Read from CMD Line

	// Read from command line

	// Connect to database

	app.Domain = "example.com"

	log.Println("Server started on port ", port)

	// Start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}

}
