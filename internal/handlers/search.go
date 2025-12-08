package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie_tracker/internal/api"
)

// Suggestion représente une suggestion de recherche
type Suggestion struct {
	Value  string `json:"value"`
	Type   string `json:"type"`   // "artist", "member", "location", "creation-date"
	Artist string `json:"artist"` // Nom de l'artiste associé
	ID     int    `json:"id"`     // ID de l'artiste pour redirection
}

// SearchSuggestions gère la barre de recherche avec suggestions
func SearchSuggestions(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("q")))

	log.Printf("🔍 Recherche: '%s'", query)

	// Si moins de 2 caractères, retourne un tableau vide
	if len(query) < 2 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]Suggestion{})
		return
	}

	// Récupère tous les artistes
	artists, err := api.GetArtists()
	if err != nil {
		log.Printf("❌ Erreur GetArtists: %v", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	var suggestions []Suggestion
	seen := make(map[string]bool) // Pour éviter les doublons

	for _, a := range artists {
		// 1️⃣ Recherche dans le nom de l'artiste
		if strings.Contains(strings.ToLower(a.Name), query) {
			key := "artist:" + a.Name
			if !seen[key] {
				suggestions = append(suggestions, Suggestion{
					Value:  a.Name,
					Type:   "artist",
					Artist: a.Name,
					ID:     a.ID,
				})
				seen[key] = true
			}
		}

		// 2️⃣ Recherche dans les membres
		for _, member := range a.Members {
			if strings.Contains(strings.ToLower(member), query) {
				key := "member:" + member
				if !seen[key] {
					suggestions = append(suggestions, Suggestion{
						Value:  member,
						Type:   "member",
						Artist: a.Name,
						ID:     a.ID,
					})
					seen[key] = true
				}
			}
		}

		// 3️⃣ Recherche dans la date de création
		yearStr := strconv.Itoa(a.CreationDate)
		if strings.Contains(yearStr, query) {
			key := "creation-date:" + yearStr + ":" + a.Name
			if !seen[key] {
				suggestions = append(suggestions, Suggestion{
					Value:  yearStr,
					Type:   "creation-date",
					Artist: a.Name,
					ID:     a.ID,
				})
				seen[key] = true
			}
		}

		// 4️⃣ Recherche dans le premier album
		if strings.Contains(strings.ToLower(a.FirstAlbum), query) {
			key := "first-album:" + a.FirstAlbum + ":" + a.Name
			if !seen[key] {
				suggestions = append(suggestions, Suggestion{
					Value:  a.FirstAlbum,
					Type:   "first-album",
					Artist: a.Name,
					ID:     a.ID,
				})
				seen[key] = true
			}
		}

		// 5️⃣ Recherche dans les locations (optimisé)
		rel, err := api.GetRelationByID(a.ID)
		if err == nil {
			for loc := range rel.DatesLocations {
				locClean := strings.ReplaceAll(strings.ToLower(loc), "_", " ")
				locClean = strings.ReplaceAll(locClean, "-", ", ")

				if strings.Contains(locClean, query) {
					key := "location:" + loc
					if !seen[key] {
						suggestions = append(suggestions, Suggestion{
							Value:  loc,
							Type:   "location",
							Artist: a.Name,
							ID:     a.ID,
						})
						seen[key] = true
						break // Une seule location par artiste
					}
				}
			}
		}
	}

	// Limite à 10 résultats
	if len(suggestions) > 10 {
		suggestions = suggestions[:10]
	}

	log.Printf("✅ Suggestions trouvées: %d", len(suggestions))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}
