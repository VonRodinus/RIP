package handlers

import (
	"RIP/internal/models"
	"html/template"
	"net/http"
	"path/filepath"
)

// CalcHandler обрабатывает страницу расчета
func CalcHandler(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Artifacts      []models.Artifact
		CurrentRequest models.CalculationRequest
	}{
		Artifacts:      models.Artifacts,
		CurrentRequest: models.CurrentRequest,
	}

	renderTemplate(w, "calculation.html", data)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("internal", "ui", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template not found: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
