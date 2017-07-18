package filterOverview

import "github.com/ONSdigital/dp-frontend-models/model"

// Page ...
type Page struct {
	model.Page
	Data     FilterOverview `json:"data"`
	FilterID string         `json:"filter_id"`
}

// FilterOverview ...
type FilterOverview struct {
	Dimensions         []Dimension `json:"dimensions"`
	PreviewAndDownload Link        `json:"preview_and_download"`
	Cancel             Link        `json:"cancel"`
}

// Dimension ...
type Dimension struct {
	Filter          string   `json:"filter"`
	AddedCategories []string `json:"added_categories"`
	Link            Link     `json:"link"`
}

// Link ...
type Link struct {
	URL   string `json:"url"`
	Label string `json:"label"`
}
