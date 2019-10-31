package item

type RequestBody struct {
	MenuItemUuid   string `json:"menuItemUuid"`
	SectionUuid    string `json:"sectionUuid"`
	SubsectionUuid string `json:"subsectionUuid"`
	StoreUuid      string `json:"storeUuid"`
}

type Response struct {
	Status string `json:"status"`
	Data   Item   `json:"data"`
}

type Item struct {
	Uuid               string          `json:"uuid"`
	Title              string          `json:"title"`
	ItemDescription    string          `json:"itemDescription"`
	Price              int             `json:"price"`
	ImageUrl           string          `json:"imageUrl"`
	CustomizationsList []Customization `json:"customizationsList"`
}

type Customization struct {
	Uuid         string   `json:"uuid"`
	Title        string   `json:"title"`
	MinPermitted int      `json:"minPermitted"`
	MaxPermitted int      `json:"maxPermitted"`
	DisplayState string   `json:"displayState"`
	Options      []Option `json:"options"`
}

type Option struct {
	Uuid            string `json:"uuid"`
	Title           string `json:"title"`
	Price           int    `json:"price"`
	DefaultQuantity int    `json:"defaultQuantity"`
	MinPermitted    int    `json:"minPermitted"`
	MaxPermitted    int    `json:"maxPermitted"`
	IsSoldOut       bool   `json:"isSoldOut"`
}
