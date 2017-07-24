package previewPage

import "github.com/ONSdigital/dp-frontend-models/model"

// Page ...
type Page struct {
	model.Page
	Data PreviewPage `json:"data"`
}

// PreviewPage ...
type PreviewPage struct {
	FilterID           string      `json:"filter_id"`
	Downloads          []Download  `json:"downloads"`
	Dimensions         []Dimension `json:"dimensions"`
	DatasetDescription string      `json:"dataset_description"`
	DatasetTitle       string      `json:"dataset_title"`
}

// Download has the details for an individual downloadable files
type Download struct {
	Extension string `json:"extension"`
	Size      string `json:"size"`
	URI       string `json:"uri"`
}

// Dimension ...
type Dimension struct {
	Title  string   `json:"title"`
	Values []string `json:"values"`
}
