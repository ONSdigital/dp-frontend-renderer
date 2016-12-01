package homepage

import (
	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/dp-frontend-models/model/dd"
)

// Homepage contains model data for the Data Discovery homepage
type Homepage struct {
	model.Page
	Datasets *dd.Datasets `json:"datasets"`
}
