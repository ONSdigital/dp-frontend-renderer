package list

import "github.com/ONSdigital/dp-frontend-models/model"

// Page represents the template data structure used for the geography list page
type Page struct {
	model.Page
	Data GeographyListPage `json:"data"`
}

// GeographyListPage represents the data on the Geography Landing page
type GeographyListPage struct {
	Items []Item `json:"items"`
}

// Item represents the child item of a geography code list
type Item struct {
	Label string `json:"label"`
	ID    string `json:"id"`
	URI   string `json:"uri"`
}
