package handlers

import (
	"RIP/internal/models"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.NotFound(w, r)
		return
	}
	orderID := pathParts[2]

	if models.CurrentRequest.ID != orderID {
		http.NotFound(w, r)
		return
	}

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
