package ageSelectorList

import "github.com/ONSdigital/dp-frontend-models/model"

// Page ...
type Page struct {
	model.Page
	Data     AgeSelectorList `json:"data"`
	FilterID string          `json:"filter_id"`
}

// AgeSelectorList ...
type AgeSelectorList struct {
	AddFromRange  Link     `json:"add_from_range"`
	SaveAndReturn Link     `json:"save_and_return"`
	Cancel        Link     `json:"cancel"`
	FiltersAmount int      `json:"filters_amount"`
	FiltersAdded  []Filter `json:"filters_added"`
	RemoveAll     Link     `json:"remove_all"`
	AgeRange      Range    `json:"age_range"`
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
	StartNum int `json:"start_num"`
	EndNum   int `json:"end_num"`
}
