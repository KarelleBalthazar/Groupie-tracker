package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie_tracker/internal/models"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

// GetArtists récupère tous les artistes
func GetArtists() ([]models.Artist, error) {
	resp, err := http.Get(baseURL + "/artists")
	if err != nil {
		return nil, fmt.Errorf("requête API artists échouée: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API artists: status %d", resp.StatusCode)
	}

	var artists []models.Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, fmt.Errorf("decode JSON artists: %w", err)
	}

	return artists, nil
}

// GetArtistByID récupère un artiste précis depuis la liste
func GetArtistByID(id int) (models.Artist, error) {
	artists, err := GetArtists()
	if err != nil {
		return models.Artist{}, err
	}

	for _, a := range artists {
		if a.ID == id {
			return a, nil
		}
	}
	return models.Artist{}, fmt.Errorf("artist %d not found", id)
}

// GetRelationByID récupère les concerts (lieux + dates) d'un artiste
// (via l'endpoint /relation/{id} déjà fourni par l'API)
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

// =====================
// DATES
// =====================

// GetDates récupère toutes les dates pour tous les artistes
func GetDates() ([]models.DateItem, error) {
	resp, err := http.Get(baseURL + "/dates")
	if err != nil {
		return nil, fmt.Errorf("requête API dates échouée: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API dates: status %d", resp.StatusCode)
	}

	var dr models.DatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&dr); err != nil {
		return nil, fmt.Errorf("decode JSON dates: %w", err)
	}

	return dr.Index, nil
}

// GetDatesByID récupère les dates d'un artiste précis
func GetDatesByID(id int) (models.DateItem, error) {
	items, err := GetDates()
	if err != nil {
		return models.DateItem{}, err
	}

	for _, it := range items {
		if it.ID == id {
			return it, nil
		}
	}

	return models.DateItem{}, fmt.Errorf("dates pour l'artiste %d introuvables", id)
}

// =====================
// LOCATIONS
// =====================

// GetLocations récupère tous les lieux pour tous les artistes
func GetLocations() ([]models.LocationItem, error) {
	resp, err := http.Get(baseURL + "/locations")
	if err != nil {
		return nil, fmt.Errorf("requête API locations échouée: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API locations: status %d", resp.StatusCode)
	}

	var lr models.LocationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&lr); err != nil {
		return nil, fmt.Errorf("decode JSON locations: %w", err)
	}

	return lr.Index, nil
}

// GetLocationByID récupère les lieux d'un artiste précis
func GetLocationByID(id int) (models.LocationItem, error) {
	items, err := GetLocations()
	if err != nil {
		return models.LocationItem{}, err
	}

	for _, it := range items {
		if it.ID == id {
			return it, nil
		}
	}

	return models.LocationItem{}, fmt.Errorf("locations pour l'artiste %d introuvables", id)
}
