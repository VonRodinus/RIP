package models

type TPQRequestItem struct {
	ArtifactID string
	Comment    string // Описание находки (из м-м)
}

type TPQCalculationRequest struct {
	ID         string
	Excavation string // Название раскопки
	TPQItems   []TPQRequestItem
	Result     string
}

var CurrentTPQRequest = TPQCalculationRequest{
	ID:         "req_001",
	Excavation: "",
	TPQItems:   []TPQRequestItem{},
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
