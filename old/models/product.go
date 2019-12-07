package models

type Product struct {
	Name   string   `json:"name"`
	Price  string   `json:"price"`
	//Type   string   `json:"type"`
	Color  string   `json:"color"`
	//Gender string   `json:"gender"`
	//Shop   string   `json:"shop"`
	Image  []string `json:"image_url"`
}
