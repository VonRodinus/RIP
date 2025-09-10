package models

// Модель артефакта (услуги)
type Artifact struct {
	ID          string
	Name        string
	Period      string
	Description string
	ImageURL    string // URL из Minio
	TPQ         int
}

// Коллекция артефактов (услуг)
var Artifacts = []Artifact{
	{
		ID:          "denarii",
		Name:        "Денарий императора Траяна",
		Period:      "98-117 гг. н.э.",
		Description: "Серебряная монета, отчеканенная во времена правления императора Траяна. Использовалась по всей Римской империи как основное средство платежа.",
		ImageURL:    "http://localhost:9000/artifacts/denarii.webp",
		TPQ:         117,
	},
	{
		ID:          "dressel20",
		Name:        "Обломок амфоры типа Dressel 20",
		Period:      "90-200 гг. н.э.",
		Description: "Фрагмент испанской оливково-масляной амфоры, широко распространенной в римский период. Характерная форма позволяет точно датировать находку.",
		ImageURL:    "http://localhost:9000/artifacts/amfora.webp",
		TPQ:         200,
	},
	{
		ID:          "girya",
		Name:        "Бронзовая гирька от римских весов",
		Period:      "120-190 гг. н.э.",
		Description: "Точная гирька для весов, использовавшаяся в римской торговле. Имеет стандартизированный вес и клеймо производителя.",
		ImageURL:    "http://localhost:9000/artifacts/girya.png",
		TPQ:         190,
	},
	{
		ID:          "dirhem",
		Name:        "Дирхем Аббасидского Халифата",
		Period:      "766-809 гг. н.э.",
		Description: "Серебряная монета Аббасидского халифата с арабскими надписями. Широко использовалась в международной торговле.",
		ImageURL:    "http://localhost:9000/artifacts/dirhem.jpg",
		TPQ:         809,
	},
}
