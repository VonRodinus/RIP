package models

type RequestItem struct {
	ArtifactID string
	Comment    string // Описание находки (из м-м)
}

type CalculationRequest struct {
	ID         string
	Excavation string // Название раскопки
	Items      []RequestItem
	Result     string
}

var CurrentRequest = CalculationRequest{
	ID:         "req_001",
	Excavation: "",
	Items:      []RequestItem{},
	Result:     "—",
}

func FindArtifactByID(id string) *Artifact {
	for _, artifact := range Artifacts {
		if artifact.ID == id {
			return &artifact
		}
	}
	return nil
}
