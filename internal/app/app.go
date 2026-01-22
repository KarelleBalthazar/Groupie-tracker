package app

import (
	"fmt"
	"net/http"
)

func Init() error {
	fmt.Println("🎵 Démarrage du serveur Groupie Tracker...")

	if err := initTemplates(); err != nil {
		return fmt.Errorf("chargement templates: %v", err)
	}

	if err := loadArtists(); err != nil {
		return fmt.Errorf("chargement artistes: %v", err)
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/artist/", artistHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("✅ Serveur prêt sur http://localhost:8080")
	return http.ListenAndServe(":8080", nil)
}
