package handlers

import (
	"RIP/internal/models"
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

func BuildingTPQCalcHandler(w http.ResponseWriter, r *http.Request) {

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.NotFound(w, r)
		return
	}
	orderID := pathParts[2]

	if models.CurrentTPQRequest.ID != orderID {
		http.NotFound(w, r)
		return
	}

	data := struct {
		Artifacts         []models.Artifact
		CurrentTPQRequest models.TPQCalculationRequest
	}{
		Artifacts:         models.Artifacts,
		CurrentTPQRequest: models.CurrentTPQRequest,
	}

	renderTemplate(w, "building_tpq_calc.html", data)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("internal", "ui", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template not found: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Создаём буфер и рендерим в него
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Если всё прошло – только теперь пишем в ответ
	w.WriteHeader(http.StatusOK)
	buf.WriteTo(w)
}
