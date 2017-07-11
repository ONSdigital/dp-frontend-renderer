package previewPage

import "github.com/ONSdigital/dp-frontend-models/model"

// Page ...
type Page struct {
	model.Page
	Data PreviewPage `json:"data"`
}

// PreviewPage ...
type PreviewPage struct {
	JobID string `json:"job_id"`
}
