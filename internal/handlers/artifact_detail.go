package handlers

import (
	"RIP/internal/models"
	"net/http"
	"strings"
)

// DetailHandler обрабатывает страницу детального просмотра артефакта
func ArtifactDetailHandler(w http.ResponseWriter, r *http.Request) {

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.NotFound(w, r)
		return
	}

	artifactID := pathParts[2]

	artifact := models.FindArtifactByID(artifactID)
	if artifact == nil {
		http.NotFound(w, r)
		return
	}

	data := struct {
		Artifact     models.Artifact
		RequestCount int
	}{
		Artifact:     *artifact,
		RequestCount: len(models.CurrentTPQRequest.TPQItems),
	}

	renderTemplate(w, "artifact-detail.html", data)
}
