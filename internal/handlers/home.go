package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"groupie_tracker/internal/api"
	"groupie_tracker/internal/models"
)

type HomePageData struct {
	Artists    []models.Artist
	NameFilter string
	YearMin    string
	YearMax    string
	AlbumMin   string
	AlbumMax   string
	MembersMin string
	MembersMax string
	Location   string
	SoloOnly   bool
}

func Home(w http.ResponseWriter, r *http.Request) {
	// Récupère tous les artistes depuis l'API
	artists, err := api.GetArtists()
	if err != nil {
		log.Printf("Erreur API artists: %v", err)
		ServerError(w, r)
		return
	}

	// Lecture des filtres envoyés par l'utilisateur
	q := r.URL.Query()
	nameFilter := strings.ToLower(strings.TrimSpace(q.Get("name")))
	yearMinStr := strings.TrimSpace(q.Get("year_min"))
	yearMaxStr := strings.TrimSpace(q.Get("year_max"))
	albumMinStr := strings.TrimSpace(q.Get("album_min"))
	albumMaxStr := strings.TrimSpace(q.Get("album_max"))
	membersMinStr := strings.TrimSpace(q.Get("members_min"))
	membersMaxStr := strings.TrimSpace(q.Get("members_max"))
	locationFilter := strings.ToLower(strings.TrimSpace(q.Get("location")))
	soloOnly := q.Get("solo_only") == "true"

	// Conversion en entiers
	yearMin, _ := strconv.Atoi(yearMinStr)
	yearMax, _ := strconv.Atoi(yearMaxStr)
	membersMin, _ := strconv.Atoi(membersMinStr)
	membersMax, _ := strconv.Atoi(membersMaxStr)

	// Liste filtrée
	var filtered []models.Artist

	for _, a := range artists {
		// Filtre par nom
		if nameFilter != "" && !strings.Contains(strings.ToLower(a.Name), nameFilter) {
			continue
		}

		// Filtre par année de création
		if yearMin > 0 && a.CreationDate < yearMin {
			continue
		}
		if yearMax > 0 && a.CreationDate > yearMax {
			continue
		}

		// Filtre par date du premier album
		if albumMinStr != "" && a.FirstAlbum < albumMinStr {
			continue
		}
		if albumMaxStr != "" && a.FirstAlbum > albumMaxStr {
			continue
		}

		// Filtre par nombre de membres
		memberCount := len(a.Members)
		if membersMin > 0 && memberCount < membersMin {
			continue
		}
		if membersMax > 0 && memberCount > membersMax {
			continue
		}

		// Filtre solo uniquement
		if soloOnly && memberCount > 1 {
			continue
		}

		// Filtre par localisation
		if locationFilter != "" {
			rel, err := api.GetRelationByID(a.ID)
			if err != nil {
				log.Printf("Erreur GetRelationByID pour artiste %d : %v", a.ID, err)
				continue
			}

			found := false
			for loc := range rel.DatesLocations {
				if strings.Contains(strings.ToLower(loc), locationFilter) {
					found = true
					break
				}
			}

			if !found {
				continue
			}
		}

		// Si tous les filtres passent → on garde l'artiste
		filtered = append(filtered, a)
	}

	// Données envoyées à la page HTML
	data := HomePageData{
		Artists:    filtered,
		NameFilter: q.Get("name"),
		YearMin:    yearMinStr,
		YearMax:    yearMaxStr,
		AlbumMin:   albumMinStr,
		AlbumMax:   albumMaxStr,
		MembersMin: membersMinStr,
		MembersMax: membersMaxStr,
		Location:   q.Get("location"),
		SoloOnly:   soloOnly,
	}

	// Rendu HTML
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Printf("Erreur parsing template : %v", err)
		ServerError(w, r)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Erreur template home.html : %v", err)
		ServerError(w, r)
	}
}
