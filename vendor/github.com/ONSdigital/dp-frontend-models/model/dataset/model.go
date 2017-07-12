package dataset

import "github.com/ONSdigital/dp-frontend-models/model"

//Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	model.Page
	Dataset Dataset `json:"data"`
}

//Dataset...
type Dataset struct {
	Foo string `json:"_"`
}
