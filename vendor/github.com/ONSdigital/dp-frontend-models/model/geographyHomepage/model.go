package geographyHomepage

import "github.com/ONSdigital/dp-frontend-models/model"

// Page contains the data re-used on each page as well as the data for the current page
type Page struct {
	model.Page
	Data GeographyList `json:"data"`
}

// GeographyList represents the data on the Geography list page
type GeographyList struct {
	Items []GeographyDataType `json:"items"`
}

// GeographyDataType represents the Type of data of the Geography list page
type GeographyDataType struct {
	Label string `json:"label"`
	ID    string `json:"id"`
}
