package main

import (
	"log"
	"net/http"
	"os"

	"groupie_tracker/internal/handlers"
	"groupie_tracker/internal/routes"
)

func main() {
	// Chargement des templates
	log.Println("📦 Chargement des templates...")
	if err := handlers.InitTemplates(); err != nil {
		log.Fatalf("❌ Erreur templates: %v", err)
	}
	log.Println("✅ Templates chargés")

	// Configuration des routes
	mux := routes.SetupRoutes()

	// Démarrage du serveur
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Printf("🚀 Serveur démarré sur http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("❌ Erreur serveur: %v", err)
	}
}
