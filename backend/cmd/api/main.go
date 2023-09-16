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
	app.Domain = "example.com"

	// Read from command line

	// Connect to database

	// Start a web server
	log.Println("Server started on port ", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		log.Fatal(err)
	}

}
