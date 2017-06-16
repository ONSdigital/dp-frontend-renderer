package middlePage

import "github.com/ONSdigital/dp-frontend-models/model"

// Page ...
type Page struct {
	model.Page
	Data MiddlePage `json:"data"`
}

// MiddlePage ...
type MiddlePage struct {
	JobID string `json:"job_id"`
}
