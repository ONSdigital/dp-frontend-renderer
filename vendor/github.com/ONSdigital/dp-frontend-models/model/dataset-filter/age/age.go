package age

import "github.com/ONSdigital/dp-frontend-models/model"

// Page represents an age selection page
type Page struct {
	model.Page
	Data     Data   `json:"data"`
	FilterID string `json:"filter_id"`
}

// Data represents the data for the age page
type Data struct {
	Youngest      string  `json:"youngest"`
	Oldest        string  `json:"oldest"`
	FirstSelected string  `json:"first_selected"`
	LastSelected  string  `json:"last_selected"`
	Ages          []Value `json:"ages"`
	CheckedRadio  string  `json:"checked_radio"`
	FormAction    Link    `json:"form_action"`
	HasAllAges    bool    `json:"has_all_ages"`
	AllAgesOption string  `json:"all_ages_option"`
}

// Value represents a single age value
type Value struct {
	Label      string
	Option     string
	IsSelected bool
}

// Link represents a link
type Link struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}
