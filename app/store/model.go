package store

type Store struct {
	Uuid     string    `json:"uuid"`
	Title    string    `json:"title"`
	Sections []Section `json:"sections"`
}

type Section struct {
	Uuid        string       `json:"uuid"`
	Title       string       `json:"title"`
	Subtitle    string       `json:"subtitle"`
	IsOnSale    bool         `json:"isOnSale"`
	Subsections []Subsection `json:"subsections"`
}

type Subsection struct {
	Uuid     string `json:"uuid"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Items    []Item `json:"items"`
}

type Item struct {
	Uuid        string `json:"uuid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageUrl    string `json:"imageUrl"`
}
