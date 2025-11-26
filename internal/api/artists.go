package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie_tracker/internal/models"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

// GetArtists va récupérer tous les artistes depuis l'API externe
func GetArtists() ([]models.Artist, error) {
	resp, err := http.Get(baseURL + "/artists")
	if err != nil {
		return nil, fmt.Errorf("requête API artists échouée: %w", err)
	}
	defer resp.Body.Close()

	// On vérifie le code HTTP
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API artists: status %d", resp.StatusCode)
	}

	var artists []models.Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, fmt.Errorf("decode JSON artists: %w", err)
	}

	return artists, nil
}

// GetRelationByID récupère les relations (dates + lieux) d'un artiste
func GetRelationByID(id int) (models.Relation, error) {
	resp, err := http.Get(fmt.Sprintf("%s/relation/%d", baseURL, id))
	if err != nil {
		return models.Relation{}, fmt.Errorf("requête API relation échouée: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Relation{}, fmt.Errorf("API relation: status %d", resp.StatusCode)
	}

	var rel models.Relation
	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		return models.Relation{}, fmt.Errorf("decode JSON relation: %w", err)
	}

	return rel, nil
}
