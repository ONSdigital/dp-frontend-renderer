package geographyListPage

import "github.com/ONSdigital/dp-frontend-models/model"

// Page contains the data re-used on each page as well as the data for the current page
type Page struct {
	model.Page
	Data GeographyListPage `json:"data"`
}

// GeographyListPage represents the Geography data on the Geography list page
type GeographyListPage struct {
	AreaTypes []AreaType `json:"items"`
}

// AreaType represents the Type of data of the Geography page
type AreaType struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
