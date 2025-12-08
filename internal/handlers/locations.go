package handlers

import (
	"net/http"
	"strings"

	"groupie_tracker/internal/api"
	"groupie_tracker/internal/models"
)

func FilterByLocation(w http.ResponseWriter, r *http.Request) {
	location := strings.ToLower(r.URL.Query().Get("location"))

	artists, _ := api.GetArtists()
	var filtered []models.Artist

	for _, a := range artists {
		rel, err := api.GetRelationByID(a.ID)
		if err != nil {
			continue
		}

		// Cherche dans les clés de DatesLocations
		for loc := range rel.DatesLocations {
			if strings.Contains(strings.ToLower(loc), location) {
				filtered = append(filtered, a)
				break
			}
		}
	}

	data := HomePageData{Artists: filtered}
	templates.ExecuteTemplate(w, "home.html", data)
}
