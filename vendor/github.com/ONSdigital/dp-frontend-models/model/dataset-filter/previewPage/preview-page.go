package previewPage

import "github.com/ONSdigital/dp-frontend-models/model"

// Page ...
type Page struct {
	model.Page
	Data             PreviewPage `json:"data"`
	IsPreviewLoaded  bool        `json:"is_preview_loaded"`
	IsDownloadLoaded bool        `json:"is_download_loaded"`
	NoDimensionData  bool        `json:"no_dimension_data"`
}

// PreviewPage ...
type PreviewPage struct {
	FilterID              string        `json:"filter_id"`
	Downloads             []Download    `json:"downloads"`
	Dimensions            []Dimension   `json:"dimensions"`
	IsLatestVersion       bool          `json:"is_latest_version"`
	LatestVersion         LatestVersion `json:"latest_version"`
	DatasetTitle          string        `json:"dataset_title"`
	DatasetID             string        `json:"dataset_id"`
	Edition               string        `json:"edition"`
	ReleaseDate           string        `json:"release_date"`
	UnitOfMeasurement     string        `json:"unit_of_measurement"`
	SingleValueDimensions []Dimension   `json:"single_value_dimensions"`
	FilterOutputID        string        `json:"filter_output_id"`
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
