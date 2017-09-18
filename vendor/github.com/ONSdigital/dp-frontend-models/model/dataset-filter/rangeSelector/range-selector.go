package rangeSelector

import "github.com/ONSdigital/dp-frontend-models/model"

// Page represents a range selector page
type Page struct {
	model.Page
	Data     RangeSelector `json:"data"`
	FilterID string        `json:"job_id"`
}

// RangeSelector represents data fields on the page
type RangeSelector struct {
	AddAllChecked     bool      `json:"add_all_checked"`
	AddAllInRange     Link      `json:"add_all"`
	AddFromList       Link      `json:"add_from_list"`
	AddRange          Link      `json:"add_range"`
	AddNewRange       Link      `json:"add_new_range"`
	Cancel            Link      `json:"cancel"`
	DateRangeData     DateRange `json:"date_range,omitempty"`
	FiltersAmount     int       `json:"filters_amount"`
	FiltersAdded      []Filter  `json:"filters_added"`
	NumberOfSelectors int       `json:"num_of_selectors"`
	RemoveRange       Link      `json:"remove_range"`
	RemoveAll         Link      `json:"remove_all"`
	RangeData         Range     `json:"range_values"`
	SaveAndReturn     Link      `json:"save_and_return"`
	Title             string    `json:"title"`
}

// Link represents a link
type Link struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

// Filter represents an item in the filter basket
type Filter struct {
	Label     string `json:"label"`
	RemoveURL string `json:"remove_url"`
	ID        string `json:"id"`
}

// DateRange extends range with month and year values
type DateRange struct {
	Range
	MonthValues []string `json:"month_values"`
	YearValues  []string `json:"year_values"`
}

// Range represents a range on the page
type Range struct {
	URL        string   `json:"url"`
	Values     []string `json:"values"`
	StartLabel string   `json:"start_label"`
	EndLabel   string   `json:"end_label"`
}
