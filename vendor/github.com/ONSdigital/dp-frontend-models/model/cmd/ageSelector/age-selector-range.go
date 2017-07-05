package ageSelectorRange

import "github.com/ONSdigital/dp-frontend-models/model"

// Page ...
type Page struct {
	model.Page
	Data  AgeSelectorRange `json:"data"`
	JobID string           `json:"job_id"`
}

// AgeSelectorRange ...
type AgeSelectorRange struct {
	AddFromList       Link     `json:"add_from_list"`
	NumberOfSelectors int      `json:"num_of_selectors"`
	AddAges           Link     `json:"add_ages"`
	AddNewRange       Link     `json:"add_new_range"`
	RemoveRange       Link     `json:"remove_range"`
	SaveAndReturn     Link     `json:"save_and_return"`
	Cancel            Link     `json:"cancel"`
	FiltersAmount     int      `json:"filters_amount"`
	FiltersAdded      []Filter `json:"filters_added"`
	RemoveAll         Link     `json:"remove_all"`
	AgeRange          Range    `json:"age_range"`
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
