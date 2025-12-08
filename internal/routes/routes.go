package routes

import (
	"net/http"

	"groupie_tracker/internal/handlers"
)

// SetupRoutes configure toutes les routes de l'application
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Fichiers statiques
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Pages principales
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/artist/", handlers.Artist)

	// Recherche (AJOUTÉ !)
	mux.HandleFunc("/search", handlers.SearchSuggestions)

	// Filtres
	mux.HandleFunc("/filter/location", handlers.FilterByLocation)

	// Pages d'erreur
	mux.HandleFunc("/404", handlers.NotFound)
	mux.HandleFunc("/500", handlers.ServerError)
	mux.HandleFunc("/400", handlers.BadRequest)

	return mux
}
