package handlers

import (
	"RIP/internal/db"
	"RIP/internal/models"
	"RIP/internal/session"
	"net/http"
	"strings"
)

// ArtifactDetailHandler godoc
// @Summary Display artifact details
// @Description Render HTML page with artifact details
// @Tags artifacts
// @Produce html
// @Param id path string true "Artifact ID"
// @Success 200 {string} string "HTML page"
// @Failure 404 {string} string "Artifact not found"
// @Router /artifact/{id} [get]
func ArtifactDetailHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.NotFound(w, r)
		return
	}
	id := pathParts[2]

	var artifact models.Artifact
	if err := db.DB.Where("id = ? AND status = ?", id, "active").First(&artifact).Error; err != nil {
		http.NotFound(w, r)
		return
	}

	sess := session.GetUser(r)
	var currentReq *models.TPQRequest
	var requestCount int
	if sess != nil {
		currentReq = getCurrentDraftRequest(sess.UserID)
		if currentReq != nil {
			requestCount = len(currentReq.TPQItems)
		}
	}

	data := struct {
		Artifact     models.Artifact
		RequestCount int
		CurrentReq   *models.TPQRequest
	}{
		Artifact:     artifact,
		RequestCount: requestCount,
		CurrentReq:   currentReq,
	}

	renderTemplate(w, "artifact_detail.html", data)
}
