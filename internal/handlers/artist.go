package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie_tracker/internal/api"
	"groupie_tracker/internal/models"
)

// ArtistPageData contient les données pour la page artiste
type ArtistPageData struct {
	Artist        models.Artist
	Relations     models.Relation
	RelationsJSON string
}

// Artist gère la page de détail d'un artiste
func Artist(w http.ResponseWriter, r *http.Request) {
	// Extraction de l'ID depuis l'URL
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		BadRequest(w, r)
		return
	}

	// Récupération des données de l'artiste
	artist, err := api.GetArtistByID(id)
	if err != nil {
		log.Printf("Erreur GetArtistByID(%d): %v", id, err)
		NotFound(w, r)
		return
	}

	// Récupération des relations (dates + lieux)
	relations, err := api.GetRelationByID(id)
	if err != nil {
		log.Printf("Erreur GetRelationByID(%d): %v", id, err)
		ServerError(w, r)
		return
	}

	// Conversion des relations en JSON pour JavaScript
	relationsJSON, err := json.Marshal(relations.DatesLocations)
	if err != nil {
		log.Printf("Erreur JSON Marshal: %v", err)
		relationsJSON = []byte("{}")
	}

	// Préparation des données
	data := ArtistPageData{
		Artist:        artist,
		Relations:     relations,
		RelationsJSON: string(relationsJSON),
	}

	// Rendu du template
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := templates.ExecuteTemplate(w, "artist.html", data); err != nil {
		log.Printf("Erreur template artist.html: %v", err)
		ServerError(w, r)
	}
}
