package geographyListPage

import "github.com/ONSdigital/dp-frontend-models/model"

// Page contains the data re-used on each page as well as the data for the current page
type Page struct {
	model.Page
	Data GeographyListPage `json:"data"`
}

// GeographyListPage represents the Geography data on the Geography list page
type GeographyListPage struct {
	Items []GeographyDataType `json:"items"`
}

// GeographyDataType represents the Type of data of the Geography page
type GeographyDataType struct {
	Label string `json:"label"`
	ID    string `json:"id"`
}
