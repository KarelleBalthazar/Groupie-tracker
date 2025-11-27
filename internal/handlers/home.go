package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie_tracker/internal/api"
	"groupie_tracker/internal/models"
)

// On charge tous les templates une seule fois pour tout le package handlers.
var templates = template.Must(template.ParseGlob("templates/*.html"))

// notFound affiche une page 404 personnalisée.
func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	if err := templates.ExecuteTemplate(w, "404.html", nil); err != nil {
		// En dernier recours, on envoie juste un texte brut
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
}

// Home gère la page d'accueil : liste de tous les artistes avec filtres.
func Home(w http.ResponseWriter, r *http.Request) {
	// On ne veut que "/"
	if r.URL.Path != "/" {
		notFound(w, r)
		return
	}

	// Récupération des artistes via l'API
	artists, err := api.GetArtists()
	if err != nil {
		log.Println("Erreur GetArtists:", err)
		http.Error(w, "Erreur lors de la récupération des artistes", http.StatusInternalServerError)
		return
	}

	// =========================
	//       LECTURE FILTRES
	// =========================
	q := r.URL.Query()

	nameFilter := strings.ToLower(strings.TrimSpace(q.Get("name")))
	yearStr := strings.TrimSpace(q.Get("year"))
	membersStr := strings.TrimSpace(q.Get("members"))
	soloOnly := q.Get("solo_only") == "on" || q.Get("solo_only") == "true"

	var yearFilter, membersCount int

	if yearStr != "" {
		if v, err := strconv.Atoi(yearStr); err == nil {
			yearFilter = v
		}
	}

	if membersStr != "" {
		if v, err := strconv.Atoi(membersStr); err == nil {
			membersCount = v
		}
	}

	// =========================
	//      APPLICATION FILTRES
	// =========================
	var filtered []models.Artist

	for _, a := range artists {
		// Filtre par nom (contient, insensible à la casse)
		if nameFilter != "" && !strings.Contains(strings.ToLower(a.Name), nameFilter) {
			continue
		}

		// Année de création exacte
		if yearFilter != 0 && a.CreationDate != yearFilter {
			continue
		}

		// Nombre de membres exact
		if membersCount != 0 && len(a.Members) != membersCount {
			continue
		}

		// Uniquement solo
		if soloOnly && len(a.Members) != 1 {
			continue
		}

		filtered = append(filtered, a)
	}

	// Si aucun filtre saisi -> on affiche tout
	noFilter := nameFilter == "" && yearFilter == 0 && membersCount == 0 && !soloOnly

	var data any
	if noFilter {
		data = artists
	} else {
		data = filtered
	}

	// Affichage avec le template home.html
	if err := templates.ExecuteTemplate(w, "home.html", data); err != nil {
		log.Println("Erreur template home:", err)
		http.Error(w, "Erreur d'affichage de la page d'accueil", http.StatusInternalServerError)
		return
	}
}
