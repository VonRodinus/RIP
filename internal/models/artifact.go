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
	{
		ID:          "glass_bracelet",
		Name:        "Обломок стеклянного браслета",
		Period:      "1100-1300 гг. н.э.",
		Description: "Серебряная монета Аббасидского халифата с арабскими надписями. Широко использовалась в международной торговле.",
		ImageURL:    "http://localhost:9000/artifacts/glass_bracelet.webp",
		TPQ:         1300,
	},
	{
		ID:          "gorshok",
		Name:        "Красноглиняный горшок",
		Period:      "1200-1400 гг. н.э.",
		Description: "Серебряная монета Аббасидского халифата с арабскими надписями. Широко использовалась в международной торговле.",
		ImageURL:    "http://localhost:9000/artifacts/gorshok.jpg",
		TPQ:         1400,
	},
	{
		ID:          "kruzhka",
		Name:        "Обломок керамической кружки",
		Period:      "1200-1450 гг. н.э.",
		Description: "Серебряная монета Аббасидского халифата с арабскими надписями. Широко использовалась в международной торговле.",
		ImageURL:    "http://localhost:9000/artifacts/kruzhka.png",
		TPQ:         1450,
	},
	{
		ID:          "stilus",
		Name:        "Костяной стилус для письма",
		Period:      "1200-1400 гг. н.э.",
		Description: "Серебряная монета Аббасидского халифата с арабскими надписями. Широко использовалась в международной торговле.",
		ImageURL:    "http://localhost:9000/artifacts/stilus.jpg",
		TPQ:         1400,
	},
	{
		ID:          "solid",
		Name:        "Солид из Любека",
		Period:      "1340-1380 гг. н.э.",
		Description: "Серебряная монета Аббасидского халифата с арабскими надписями. Широко использовалась в международной торговле.",
		ImageURL:    "http://localhost:9000/artifacts/solid.webp",
		TPQ:         1380,
	},
	{
		ID:          "glass_fragment",
		Name:        "Фрагмент оконного витража",
		Period:      "1400-1450 гг. н.э.",
		Description: "Серебряная монета Аббасидского халифата с арабскими надписями. Широко использовалась в международной торговле.",
		ImageURL:    "http://localhost:9000/artifacts/glass_fragment.webp",
		TPQ:         1450,
	},
	{
		ID:          "gulden",
		Name:        "Виттен из Любека",
		Period:      "1350-1420 гг. н.э.",
		Description: "Серебряная монета Аббасидского халифата с арабскими надписями. Широко использовалась в международной торговле.",
		ImageURL:    "http://localhost:9000/artifacts/gulden.jpg",
		TPQ:         1420,
	},
	{
		ID:          "zastezhka",
		Name:        "Застёжка рыцарского пояса",
		Period:      "1440-1480 гг. н.э.",
		Description: "Серебряная монета Аббасидского халифата с арабскими надписями. Широко использовалась в международной торговле.",
		ImageURL:    "http://localhost:9000/artifacts/zastezhka.jpg",
		TPQ:         1480,
	},
	{
		ID:          "cannon_folders",
		Name:        "Каменные пушечные ядра малого калибра",
		Period:      "1400-1470 гг. н.э.",
		Description: "Серебряная монета Аббасидского халифата с арабскими надписями. Широко использовалась в международной торговле.",
		ImageURL:    "http://localhost:9000/artifacts/cannon_folders.webp",
		TPQ:         1470,
	},
}
