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
	FilterID        string        `json:"filter_id"`
	Downloads       []Download    `json:"downloads"`
	Dimensions      []Dimension   `json:"dimensions"`
	IsLatestVersion bool          `json:"is_latest_version"`
	LatestVersion   LatestVersion `json:"latest_version"`
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

// LatestVersion ...
type LatestVersion struct {
	DatasetLandingPageURL          string `json:"dataset_landing_page_url"`
	FilterJourneyWithLatestJourney string `json:"filter_journey_with_latest_version"`
}
