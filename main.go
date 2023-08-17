package main

import (
	"log"
	"net/http"
	"todolist/models"
	"todolist/routes"
)

func main() {

	err := models.InitDB()
	if err != nil {
		panic("Failed to connect to the database")
	}
	defer models.CloseDB()

	// Routes
	r := routes.Routes()

	// Server
	log.Println("Starting server on port: 5000")
	err = http.ListenAndServe(":5000", r)
	if err != nil {
		log.Fatal(err)
	}
}
