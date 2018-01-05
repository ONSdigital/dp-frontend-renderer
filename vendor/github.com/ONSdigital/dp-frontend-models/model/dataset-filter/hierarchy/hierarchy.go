package hierarchy

import "github.com/ONSdigital/dp-frontend-models/model"

// Page ...
type Page struct {
	model.Page
	Data     Hierarchy `json:"data"`
	FilterID string    `json:"filter_id"`
}

// Hierarchy ...
type Hierarchy struct {
	Title         string   `json:"title"`
	SaveAndReturn Link     `json:"save_and_return"`
	Cancel        Link     `json:"cancel"`
	FiltersAmount string   `json:"filters_amount"`
	AddAllFilters AddAll   `json:"add_all"`
	FilterList    []List   `json:"filter_list"`
	FiltersAdded  []Filter `json:"filters_added"`
	RemoveAll     Link     `json:"remove_all"`
	GoBack        Link     `json:"go_back"`
	DimensionName string   `json:"dimension_name"`
	Parent        string   `json:"parent"`
	Type          string   `json:"type"`
	Metadata      Metadata `json:"metadata"`
}

// AddAll ...
type AddAll struct {
	Amount string `json:"amount"`
	URL    string `json:"url"`
}

// Metadata ...
type Metadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Filter ...
type Filter struct {
	Label     string `json:"label"`
	RemoveURL string `json:"remove_url"`
	ID        string `json:"id"`
}

// List ...
type List struct {
	Label    string `json:"label"`
	Selected bool   `json:"selected"`
	SubNum   string `json:"sub_num"`
	ID       string `json:"id"`
	SubType  string `json:"sub_type"`
	SubURL   string `json:"sub_url"`
}

// Link ...
type Link struct {
	URL   string `json:"url"`
	Label string `json:"label"`
}
