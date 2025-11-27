package main

import (
	"log"
	"net/http"

	"groupie_tracker/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Page d'accueil -> handlers.Home
	mux.HandleFunc("/", handlers.Home)

	// Page artiste -> handlers.Artist
	mux.HandleFunc("/artist/", handlers.Artist)

	// Fichiers statiques (JS, CSS, images...)
	static := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", static))

	log.Println("Serveur lancé sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
