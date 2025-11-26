package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie_tracker/internal/api"
	"groupie_tracker/internal/models"
)

// ArtistPageData contient les données envoyées au template artist.html
type ArtistPageData struct {
	Artist   models.Artist
	Relation models.Relation
}

// Artist gère la page /artist/{id}
func Artist(w http.ResponseWriter, r *http.Request) {
	// L'URL doit commencer par /artist/
	if !strings.HasPrefix(r.URL.Path, "/artist/") {
		http.NotFound(w, r)
		return
	}

	// On récupère l'ID dans l'URL : /artist/1 -> "1"
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Récupérer tous les artistes
	artists, err := api.GetArtists()
	if err != nil {
		log.Println("Erreur GetArtists:", err)
		http.Error(w, "Erreur lors de la récupération des artistes", http.StatusInternalServerError)
		return
	}

	// Chercher l'artiste avec le bon ID
	var artist models.Artist
	found := false
	for _, a := range artists {
		if a.ID == id {
			artist = a
			found = true
			break
		}
	}
	if !found {
		http.NotFound(w, r)
		return
	}

	// Récupérer les relations (dates + lieux) pour cet artiste
	relation, err := api.GetRelationByID(id)
	if err != nil {
		log.Println("Erreur GetRelationByID:", err)
		http.Error(w, "Erreur lors de la récupération des concerts", http.StatusInternalServerError)
		return
	}

	// Construire les données envoyées au template
	data := ArtistPageData{
		Artist:   artist,
		Relation: relation,
	}

	// templates vient déjà de home.go (variable globale du package handlers)
	if err := templates.ExecuteTemplate(w, "artist.html", data); err != nil {
		log.Println("Erreur template artist:", err)
		http.Error(w, "Erreur d'affichage de la page artiste", http.StatusInternalServerError)
		return
	}
}
