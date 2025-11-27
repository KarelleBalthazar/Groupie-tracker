package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie_tracker/internal/api"
	"groupie_tracker/internal/models"
)

type ArtistPageData struct {
	Artist   models.Artist
	Relation models.Relation
}

func Artist(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/artist/") {
		notFound(w, r)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		notFound(w, r)
		return
	}

	artist, err := api.GetArtistByID(id)
	if err != nil {
		log.Println("Artist not found:", err)
		notFound(w, r)
		return
	}

	relation, err := api.GetRelationByID(id)
	if err != nil {
		log.Println("Erreur GetRelationByID:", err)
		http.Error(w, "Erreur lors de la récupération des concerts", http.StatusInternalServerError)
		return
	}

	data := ArtistPageData{
		Artist:   artist,
		Relation: relation,
	}

	if err := templates.ExecuteTemplate(w, "artist.html", data); err != nil {
		log.Println("Erreur template artist:", err)
		http.Error(w, "Erreur d'affichage de la page artiste", http.StatusInternalServerError)
		return
	}
}
