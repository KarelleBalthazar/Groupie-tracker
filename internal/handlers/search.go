package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"groupie_tracker/internal/api"
)

type Suggestion struct {
	Value string `json:"value"`
	Type  string `json:"type"` // "artist", "member", "location", "date"
}

func SearchSuggestions(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("q"))
	if query == "" {
		json.NewEncoder(w).Encode([]Suggestion{})
		return
	}

	artists, _ := api.GetArtists()
	var suggestions []Suggestion

	for _, a := range artists {
		// Recherche dans le nom de l'artiste
		if strings.Contains(strings.ToLower(a.Name), query) {
			suggestions = append(suggestions, Suggestion{
				Value: a.Name,
				Type:  "artist",
			})
		}

		// Recherche dans les membres
		for _, member := range a.Members {
			if strings.Contains(strings.ToLower(member), query) {
				suggestions = append(suggestions, Suggestion{
					Value: member + " (" + a.Name + ")",
					Type:  "member",
				})
			}
		}

		// Recherche dans la date de création
		yearStr := strconv.Itoa(a.CreationDate)
		if strings.Contains(yearStr, query) {
			suggestions = append(suggestions, Suggestion{
				Value: a.Name + " - " + yearStr,
				Type:  "creation-date",
			})
		}

		// Recherche dans les locations
		rel, _ := api.GetRelationByID(a.ID)
		for loc := range rel.DatesLocations {
			if strings.Contains(strings.ToLower(loc), query) {
				suggestions = append(suggestions, Suggestion{
					Value: loc,
					Type:  "location",
				})
				break
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}
