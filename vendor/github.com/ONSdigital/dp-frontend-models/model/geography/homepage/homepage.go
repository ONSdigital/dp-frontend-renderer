package homepage

import "github.com/ONSdigital/dp-frontend-models/model"

// Page contains the data re-used on each page as well as the data for the current page
type Page struct {
	model.Page
	Data GeographyHomepagePage `json:"data"`
}

// GeographyHomepagePage represents the data on the geography journey's homepage
type GeographyHomepagePage struct {
	Items []Items `json:"items"`
}

// Items represents the Type of data of the Geography page
type Items struct {
	Label string `json:"label"`
	ID    string `json:"id"`
}
