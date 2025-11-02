package handlers

import (
	"RIP/internal/db"
	"RIP/internal/models"
	"RIP/internal/session"
	"net/http"
	"strings"
)

// ArtifactCatalogHandler godoc
// @Summary Display artifact catalog
// @Description Render HTML catalog of artifacts (optional auth)
// @Tags artifacts
// @Produce html
// @Param artifact_name_or_tpq_filter query string false "Filter by name or TPQ"
// @Success 200 {string} string "HTML page"
// @Failure 404 {string} string "Not found"
// @Router / [get]
func ArtifactCatalogHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	sess := session.GetUser(r)
	searchQuery := r.URL.Query().Get("artifact_name_or_tpq_filter")
	filteredArtifacts := filterArtifacts(searchQuery)

	var currentReq *models.TPQRequest
	var requestCount int
	var currentTPQRequest models.TPQRequest
	if sess != nil {
		currentReq = getCurrentDraftRequest(sess.UserID)
		if currentReq != nil {
			currentTPQRequest = *currentReq
			requestCount = len(currentReq.TPQItems)
		}
	}

	data := struct {
		Artifacts         []models.Artifact
		SearchQuery       string
		RequestCount      int
		CurrentTPQRequest models.TPQRequest
	}{
		Artifacts:         filteredArtifacts,
		SearchQuery:       searchQuery,
		RequestCount:      requestCount,
		CurrentTPQRequest: currentTPQRequest,
	}

	renderTemplate(w, "artifact_catalog.html", data)
}

func filterArtifacts(query string) []models.Artifact {
	var artifacts []models.Artifact
	q := db.DB.Where("status = ?", "active")
	if query != "" {
		searchTerm := "%" + strings.ToLower(query) + "%"
		q = q.Where("LOWER(name) LIKE ? OR start_date::text LIKE ? OR end_date::text LIKE ? OR LOWER(epoch) LIKE ?", searchTerm, searchTerm, searchTerm, searchTerm)
	}
	q.Find(&artifacts)
	return artifacts
}
