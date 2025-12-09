package handlers

import (
	"html/template"
	"log"
)

var templates *template.Template

// InitTemplates charge tous les templates HTML au démarrage
func InitTemplates() error {
	var err error
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		return err
	}

	log.Println("📄 Templates chargés:")
	for _, t := range templates.Templates() {
		log.Printf("   - %s", t.Name())
	}

	return nil
}
