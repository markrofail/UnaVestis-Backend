package models

type Product struct {
	Name  string   `json:"name"`
	Price string   `json:"price"`
	Color string   `json:"color"`
	Image []string `json:"image_url"`
	//Type   string   `json:"type"`
	//Gender string   `json:"gender"`
	//Shop   string   `json:"shop"`
}
