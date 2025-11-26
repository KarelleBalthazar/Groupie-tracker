package handlers

import (
	"html/template"
	"log"
	"net/http"

	"groupie_tracker/internal/api"
)

// On charge tous les templates une seule fois pour tout le package handlers.
var templates = template.Must(template.ParseGlob("templates/*.html"))

// notFound affiche une page 404 personnalisée.
func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	if err := templates.ExecuteTemplate(w, "404.html", nil); err != nil {
		// En dernier recours, on envoie juste un texte brut
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
}

// Home gère la page d'accueil : liste de tous les artistes.
func Home(w http.ResponseWriter, r *http.Request) {
	// On ne veut que "/"
	if r.URL.Path != "/" {
		notFound(w, r)
		return
	}

	// Récupération des artistes via l'API
	artists, err := api.GetArtists()
	if err != nil {
		log.Println("Erreur GetArtists:", err)
		http.Error(w, "Erreur lors de la récupération des artistes", http.StatusInternalServerError)
		return
	}

	// Affichage avec le template home.html
	if err := templates.ExecuteTemplate(w, "home.html", artists); err != nil {
		log.Println("Erreur template home:", err)
		http.Error(w, "Erreur d'affichage de la page d'accueil", http.StatusInternalServerError)
		return
	}
}
