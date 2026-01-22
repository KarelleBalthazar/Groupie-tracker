package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var artists []Artist
var filteredArtists []Artist

func loadArtists() error {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("erreur requête API: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("erreur lecture réponse: %v", err)
	}

	if err := json.Unmarshal(body, &artists); err != nil {
		return fmt.Errorf("erreur décodage JSON: %v", err)
	}

	// Filtrer les artistes inappropriés
	filteredArtists = []Artist{}
	for _, artist := range artists {
		// Enlever Mamonas et autres artistes avec contenu sexuel
		if !strings.Contains(strings.ToLower(artist.Name), "mamonas") &&
			!strings.Contains(strings.ToLower(artist.Image), "sex") &&
			!strings.Contains(strings.ToLower(artist.Name), "sex") {
			filteredArtists = append(filteredArtists, artist)
		}
	}

	fmt.Printf("✅ %d artistes chargés (%d après filtrage)\n", len(artists), len(filteredArtists))
	return nil
}
