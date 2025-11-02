package handlers

import (
	"RIP/internal/db"
	"RIP/internal/models"
	"net/http"
	"strings"
)

func BuildingTPQCalcHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.NotFound(w, r)
		return
	}
	orderID := pathParts[2]

	var req models.TPQRequest
	if db.DB.Preload("TPQItems.Artifact").Where("id = ? AND status != ? AND creator_id = ?", orderID, "deleted", 1).First(&req).Error != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data := struct {
		CurrentTPQRequest models.TPQRequest
	}{
		CurrentTPQRequest: req,
	}

	renderTemplate(w, "building_tpq_calc.html", data)
}
