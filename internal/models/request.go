package models

type RequestItem struct {
	ArtifactID string
	Comment    string
}

type CalculationRequest struct {
	ID     string
	Items  []RequestItem
	Result string
}

var CurrentRequest = CalculationRequest{
	ID:     "req_001",
	Items:  []RequestItem{},
	Result: "—",
}

func FindArtifactByID(id string) *Artifact {
	for _, artifact := range Artifacts {
		if artifact.ID == id {
			return &artifact
		}
	}
	return nil
}
