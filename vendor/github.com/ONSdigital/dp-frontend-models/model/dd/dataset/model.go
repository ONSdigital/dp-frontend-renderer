package dataset

import (
	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/dp-frontend-models/model/dd"
)

// Page contains model data for a dataset page
type Page struct {
	model.Page
	Dataset *dd.Dataset `json:"dataset"`
}
