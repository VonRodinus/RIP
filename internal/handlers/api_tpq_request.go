package handlers

import (
	"RIP/internal/db"
	"RIP/internal/models"
	"RIP/internal/session"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// DeleteTPQRequestItem godoc
// @Summary Delete TPQ request item
// @Description Delete an item from TPQ request (user required, own draft)
// @Tags tpq_requests
// @Param request_id path string true "Request ID"
// @Param artifact_id path string true "Artifact ID"
// @Success 204 {string} string "No Content"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Item not found"
// @Security BearerAuth
// @Router /api/tpq_requests/{request_id}/items/{artifact_id} [delete]
func DeleteTPQRequestItem(w http.ResponseWriter, r *http.Request) {
	sess := session.GetUser(r)
	if sess == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 6 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	requestID := pathParts[3]
	artifactID := pathParts[5]
	var req models.TPQRequest
	if err := db.DB.Where("id = ?", requestID).First(&req).Error; err != nil {
		http.Error(w, "Request not found", http.StatusNotFound)
		return
	}
	if req.CreatorID != sess.UserID || req.Status != "draft" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	var item models.TPQRequestItem
	if err := db.DB.Where("request_id = ? AND artifact_id = ?", requestID, artifactID).First(&item).Error; err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	if err := db.DB.Delete(&item).Error; err != nil {
		http.Error(w, "Error deleting item", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// UpdateTPQRequestItem godoc
// @Summary Update TPQ request item
// @Description Update comment on TPQ request item (user required, own draft)
// @Tags tpq_requests
// @Accept json
// @Produce json
// @Param request_id path string true "Request ID"
// @Param artifact_id path string true "Artifact ID"
// @Param item body models.TPQRequestItem true "Updated item data"
// @Success 200 {object} models.TPQRequestItem
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Item not found"
// @Security BearerAuth
// @Router /api/tpq_requests/{request_id}/items/{artifact_id} [put]
func UpdateTPQRequestItem(w http.ResponseWriter, r *http.Request) {
	sess := session.GetUser(r)
	if sess == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 6 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	requestID := pathParts[3]
	artifactID := pathParts[5]
	var req models.TPQRequest
	if err := db.DB.Where("id = ?", requestID).First(&req).Error; err != nil {
		http.Error(w, "Request not found", http.StatusNotFound)
		return
	}
	if req.CreatorID != sess.UserID || req.Status != "draft" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	var item models.TPQRequestItem
	if err := db.DB.Preload("Artifact").Where("request_id = ? AND artifact_id = ?", requestID, artifactID).First(&item).Error; err != nil {
		log.Printf("Item not found: request_id=%s, artifact_id=%s, error=%v", requestID, artifactID, err)
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	var updates models.TPQRequestItem
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		log.Printf("Invalid request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	item.Comment = updates.Comment
	if err := db.DB.Save(&item).Error; err != nil {
		log.Printf("Error updating item: %v", err)
		http.Error(w, "Error updating item", http.StatusInternalServerError)
		return
	}
	log.Printf("Updated item: request_id=%s, artifact_id=%s, comment=%s", item.RequestID, item.ArtifactID, item.Comment)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}
