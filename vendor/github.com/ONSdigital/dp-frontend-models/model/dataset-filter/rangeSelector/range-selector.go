package rangeSelector

import "github.com/ONSdigital/dp-frontend-models/model"

// Page ...
type Page struct {
	model.Page
	Data     RangeSelector `json:"data"`
	FilterID string        `json:"job_id"`
}

// RangeSelector ...
type RangeSelector struct {
	Title             string   `json:"title"`
	AddFromList       Link     `json:"add_from_list"`
	NumberOfSelectors int      `json:"num_of_selectors"`
	AddAllInRange     Link     `json:"add_all"`
	AddRange          Link     `json:"add_range"`
	AddNewRange       Link     `json:"add_new_range"`
	RemoveRange       Link     `json:"remove_range"`
	SaveAndReturn     Link     `json:"save_and_return"`
	Cancel            Link     `json:"cancel"`
	FiltersAmount     int      `json:"filters_amount"`
	FiltersAdded      []Filter `json:"filters_added"`
	RemoveAll         Link     `json:"remove_all"`
	RangeData         Range    `json:"range_values"`
}

// Link ...
type Link struct {
	URL   string `json:"url"`
	Label string `json:"label"`
}

// Filter ...
type Filter struct {
	Label     string `json:"label"`
	RemoveURL string `json:"remove_url"`
}

// Range ...
type Range struct {
	URL        string   `json:"url"`
	Values     []string `json:"values"`
	StartLabel string   `json:"start_label"`
	EndLabel   string   `json:"end_label"`
}
