package app

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var tmpl *template.Template

func initTemplates() error {
	path := filepath.Join("templates", "*.html")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		path = filepath.Join("..", "..", "templates", "*.html")
	}

	var err error
	tmpl, err = template.New("").Funcs(template.FuncMap{
		"lower": strings.ToLower,
	}).ParseGlob(path)
	return err
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page non trouvée", http.StatusNotFound)
		return
	}

	data := struct {
		Artists []Artist
		Title   string
	}{
		Artists: filteredArtists,
		Title:   "Groupie Tracker",
	}

	tmpl.ExecuteTemplate(w, "index.html", data)
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 || id >= len(filteredArtists) {
		http.Error(w, "Artiste non trouvé", http.StatusNotFound)
		return
	}

	data := struct {
		Artist Artist
		Title  string
	}{
		Artist: filteredArtists[id],
		Title:  filteredArtists[id].Name,
	}

	tmpl.ExecuteTemplate(w, "artist.html", data)
}
