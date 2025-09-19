package models

// Модель артефакта (услуги)
type Artifact struct {
	ID          string
	Name        string
	StartDate   int
	EndDate     int
	Epoch       string
	Description string
	ImageURL    string // URL из Minio
	TPQ         int
}

// Коллекция артефактов (услуг)
var Artifacts = []Artifact{
	{
		ID:          "denarii",
		Name:        "Денарий императора Траяна",
		StartDate:   98,
		EndDate:     117,
		Epoch:       "гг. н.э.",
		Description: "Серебряная монета, отчеканенная во времена правления императора Траяна. Использовалась по всей Римской империи как основное средство платежа.",
		ImageURL:    "http://localhost:9000/artifacts/denarii.webp",
		TPQ:         117,
	},
	{
		ID:          "dressel20",
		Name:        "Обломок амфоры типа Dressel 20",
		StartDate:   90,
		EndDate:     200,
		Epoch:       "гг. н.э.",
		Description: "Фрагмент испанской оливково-масляной амфоры, широко распространенной в римский период. Характерная форма позволяет точно датировать находку.",
		ImageURL:    "http://localhost:9000/artifacts/amfora.webp",
		TPQ:         200,
	},
	{
		ID:          "girya",
		Name:        "Бронзовая гирька от римских весов",
		StartDate:   120,
		EndDate:     190,
		Epoch:       "гг. н.э.",
		Description: "Точная гирька для весов, использовавшаяся в римской торговле. Имеет стандартизированный вес и клеймо производителя.",
		ImageURL:    "http://localhost:9000/artifacts/girya.png",
		TPQ:         190,
	},
	{
		ID:          "dirhem",
		Name:        "Дирхем Аббасидского Халифата",
		StartDate:   766,
		EndDate:     809,
		Epoch:       "гг. н.э.",
		Description: "Серебряная монета Аббасидского халифата с арабскими надписями. Широко использовалась в международной торговле.",
		ImageURL:    "http://localhost:9000/artifacts/dirhem.jpg",
		TPQ:         809,
	},
	{
		ID:          "glass_bracelet",
		Name:        "Обломок стеклянного браслета",
		StartDate:   1100,
		EndDate:     1300,
		Epoch:       "гг. н.э.",
		Description: "Фрагмент стеклянного браслета, характерного для древнерусских памятников. Цвет и форма скрутки являются датирующим признаком.",
		ImageURL:    "http://localhost:9000/artifacts/glass_bracelet.webp",
		TPQ:         1300,
	},
	{
		ID:          "gorshok",
		Name:        "Красноглиняный горшок",
		StartDate:   1200,
		EndDate:     1400,
		Epoch:       "гг. н.э.",
		Description: "Фрагмент лепного или кругового сосуда, характерного для бытовой керамики средневекового периода.",
		ImageURL:    "http://localhost:9000/artifacts/gorshok.jpg",
		TPQ:         1400,
	},
	{
		ID:          "kruzhka",
		Name:        "Обломок керамической кружки",
		StartDate:   1200,
		EndDate:     1450,
		Epoch:       "гг. н.э.",
		Description: "Фрагмент стенки гончарной кружки с характерной профилировкой.",
		ImageURL:    "http://localhost:9000/artifacts/kruzhka.png",
		TPQ:         1450,
	},
	{
		ID:          "stilus",
		Name:        "Костяной стилус для письма",
		StartDate:   1200,
		EndDate:     1400,
		Epoch:       "гг. н.э.",
		Description: "Письменный прибор для нанесения текста на восковые таблички.",
		ImageURL:    "http://localhost:9000/artifacts/stilus.jpg",
		TPQ:         1400,
	},
	{
		ID:          "solid",
		Name:        "Солид из Любека",
		StartDate:   1340,
		EndDate:     1380,
		Epoch:       "гг. н.э.",
		Description: "Серебряная монета, отчеканенная в вольном ганзейском городе Любек.",
		ImageURL:    "http://localhost:9000/artifacts/solid.webp",
		TPQ:         1380,
	},
	{
		ID:          "glass_fragment",
		Name:        "Фрагмент оконного витража",
		StartDate:   1400,
		EndDate:     1450,
		Epoch:       "гг. н.э.",
		Description: "Осколок цветного стекла от оконного витража. Свидетельствует о наличии богато украшенных строений, вероятно, культового или представительского назначения.",
		ImageURL:    "http://localhost:9000/artifacts/glass_fragment.webp",
		TPQ:         1450,
	},
	{
		ID:          "gulden",
		Name:        "Виттен из Любека",
		StartDate:   1350,
		EndDate:     1420,
		Epoch:       "гг. н.э.",
		Description: "Золотая монета отчеканенная в вольном ганзейском городе Любек.",
		ImageURL:    "http://localhost:9000/artifacts/gulden.jpg",
		TPQ:         1420,
	},
	{
		ID:          "zastezhka",
		Name:        "Застёжка рыцарского пояса",
		StartDate:   1440,
		EndDate:     1480,
		Epoch:       "гг. н.э.",
		Description: "Литая металлическая деталь от парадного поясного набора. Украшена орнаментом и может содержать геральдические элементы, указывающие на статус владельца.",
		ImageURL:    "http://localhost:9000/artifacts/zastezhka.jpg",
		TPQ:         1480,
	},
	{
		ID:          "cannon_folders",
		Name:        "Каменные пушечные ядра малого калибра",
		StartDate:   1400,
		EndDate:     1470,
		Epoch:       "гг. н.э.",
		Description: "Каменные ядра для ранних образцов артиллерии, распространённых в Европе (например, кулеврин).",
		ImageURL:    "http://localhost:9000/artifacts/cannon_folders.webp",
		TPQ:         1470,
	},
}
