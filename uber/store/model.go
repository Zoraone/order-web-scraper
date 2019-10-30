package store

type Response struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	Uuid  string `json:"uuid"`
	Title string `json:"title"`

	SectionEntitiesMap map[string]interface{} `json:"sectionEntitiesMap"`
	Sections           []Section              `json:"sections"`
	SubsectionsMap     map[string]Subsection  `json:"subsectionsMap"`
}

type SectionEntities struct {
	ItemsMap map[string]Item `json:"-"`
}

type Item struct {
	Uuid        string `json:"uuid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageUrl    string `json:"imageUrl"`
}

type Section struct {
	Uuid            string   `json:"uuid"`
	Title           string   `json:"title"`
	Subtitle        string   `json:"subtitle"`
	IsTop           bool     `json:"isTop"`
	IsOnSale        bool     `json:"isOnSale"`
	SubsectionUuids []string `json:"subsectionUuids"`
}

type Subsection struct {
	Uuid        string      `json:"uuid"`
	Title       string      `json:"title"`
	Subtitle    string      `json:"subtitle"`
	ItemUuids   []string    `json:"itemUuids"`
	DisplayType interface{} `json:"displayType"`
}
