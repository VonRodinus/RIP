package handlers

import (
	"RIP/internal/models"
	"fmt"
	"net/http"
)

// CatalogHandler обрабатывает главную страницу с каталогом артефактов
func ArtifactCatalogHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	searchQuery := r.URL.Query().Get("artifact_name_or_tpq_filter")

	filteredArtifacts := filterArtifacts(searchQuery)

	data := struct {
		Artifacts         []models.Artifact
		SearchQuery       string
		RequestCount      int
		CurrentTPQRequest models.TPQCalculationRequest
	}{
		Artifacts:         filteredArtifacts,
		SearchQuery:       searchQuery,
		RequestCount:      len(models.CurrentTPQRequest.TPQItems),
		CurrentTPQRequest: models.CurrentTPQRequest,
	}

	renderTemplate(w, "artifact_catalog.html", data)
}

func filterArtifacts(query string) []models.Artifact {
	if query == "" {
		return models.Artifacts
	}

	var filtered []models.Artifact
	for _, artifact := range models.Artifacts {
		if contains(artifact.Name, query) ||
			contains(fmt.Sprintf("%d", artifact.StartDate), query) ||
			contains(fmt.Sprintf("%d", artifact.EndDate), query) ||
			contains(artifact.Epoch, query) {
			filtered = append(filtered, artifact)
		}
	}
	return filtered
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
