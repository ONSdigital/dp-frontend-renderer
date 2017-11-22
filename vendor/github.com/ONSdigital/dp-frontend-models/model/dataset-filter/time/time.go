package time

import "github.com/ONSdigital/dp-frontend-models/model"

// Page represents a time selection page
type Page struct {
	model.Page
	Data     Data   `json:"data"`
	FilterID string `json:"filter_id"`
}

// Data represents the data for the time page
type Data struct {
	LatestTime         Value    `json:"latest_value"`
	FirstTime          Value    `json:"fist_time"`
	Values             []Value  `json:"values"`
	Months             []string `json:"months"`
	Years              []string `json:"years"`
	CheckedRadio       string   `json:"checked_radio"`
	FormAction         Link     `json:"form_action"`
	SelectedStartMonth string   `json:"selected_start_month"`
	SelectedStartYear  string   `json:"selected_start_year"`
	SelectedEndMonth   string   `json:"selected_end_month"`
	SelectedEndYear    string   `json:"selected_end_year"`
	Type               string   `json:"type"`
}

// Link represents a link
type Link struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

// Value represents a single time value
type Value struct {
	Month      string `json:"month,omitempty"`
	Year       string `json:"year,omitempty"`
	Option     string `json:"option"`
	IsSelected bool   `json:"is_selected"`
}
