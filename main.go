package main

import (
	"log"

	"groupie_tracker/cmd"
	"groupie_tracker/internal/handlers"
)

func main() {
	log.Println("📦 Chargement des templates...")
	if err := handlers.InitTemplates(); err != nil {
		log.Fatalf("❌ Erreur: %v", err)
	}
	log.Println("✅ Templates chargés")
	cmd.StartServer()
}
