package main

import (
	"log"
	"net/http"

	"groupie_tracker/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Routes principales
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/artist/", handlers.Artist)

	log.Println("Serveur lancé sur http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
