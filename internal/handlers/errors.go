package handlers

import (
	"log"
	"net/http"
)

// NotFound gère les erreurs 404
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	if err := templates.ExecuteTemplate(w, "404.html", nil); err != nil {
		log.Printf("❌ Erreur template 404.html: %v", err)
		http.Error(w, "404 - Page non trouvée", http.StatusNotFound)
	}
}

// ServerError gère les erreurs 500
func ServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := templates.ExecuteTemplate(w, "500.html", nil); err != nil {
		log.Printf("❌ Erreur template 500.html: %v", err)
		http.Error(w, "500 - Erreur interne du serveur", http.StatusInternalServerError)
	}
}

// BadRequest gère les erreurs 400
func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	if err := templates.ExecuteTemplate(w, "400.html", nil); err != nil {
		log.Printf("❌ Erreur template 400.html: %v", err)
		http.Error(w, "400 - Requête invalide", http.StatusBadRequest)
	}
}
