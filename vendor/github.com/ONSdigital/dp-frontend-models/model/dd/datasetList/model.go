package datasetList

import (
	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/dp-frontend-models/model/dd"
)

// DatasetList contains model data for the list of all Data Discovery datasets
type DatasetList struct {
	model.Page
	Datasets *dd.Datasets `json:"datasets"`
}
