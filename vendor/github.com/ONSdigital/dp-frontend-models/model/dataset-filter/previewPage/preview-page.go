package previewPage

import "github.com/ONSdigital/dp-frontend-models/model"

// Page ...
type Page struct {
	model.Page
	Data            PreviewPage `json:"data"`
	IsContentLoaded bool        `json:"is_content_loaded"`
}

// PreviewPage ...
type PreviewPage struct {
	FilterID   string      `json:"filter_id"`
	Downloads  []Download  `json:"downloads"`
	Dimensions []Dimension `json:"dimensions"`
}

// Download has the details for an individual downloadable files
type Download struct {
	Extension string `json:"extension"`
	Size      string `json:"size"`
	URI       string `json:"uri"`
}

// Dimension ...
type Dimension struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
