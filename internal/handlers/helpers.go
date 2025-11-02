package handlers

import (
	"RIP/internal/db"
	"RIP/internal/models"
	"html/template"
	"net/http"
)

func getCurrentDraftRequest(userID uint) *models.TPQRequest {
	var req models.TPQRequest
	err := db.DB.Preload("TPQItems").Where("status = ? AND creator_id = ?", "draft", userID).First(&req).Error
	if err != nil {
		return nil
	}
	return &req
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
