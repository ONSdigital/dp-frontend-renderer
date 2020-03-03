package cookiespreferences

import "github.com/ONSdigital/dp-frontend-models/model"

// Page model data for the cookies preferences form
type Page struct {
	model.Page
	PreferencesUpdated bool `json:"preferences_updated"`
}
