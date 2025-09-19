package handlers

import (
	"RIP/internal/models"
	"fmt"
	"net/http"
)

// CatalogHandler обрабатывает главную страницу с каталогом артефактов
func CatalogHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	searchQuery := r.URL.Query().Get("search")

	filteredArtifacts := filterArtifacts(searchQuery)

	data := struct {
		Artifacts      []models.Artifact
		SearchQuery    string
		RequestCount   int
		CurrentRequest models.CalculationRequest
	}{
		Artifacts:      filteredArtifacts,
		SearchQuery:    searchQuery,
		RequestCount:   len(models.CurrentRequest.Items),
		CurrentRequest: models.CurrentRequest,
	}

	renderTemplate(w, "index.html", data)
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
