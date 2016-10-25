package productPage

import "github.com/ONSdigital/dp-frontend-renderer/model"

//Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	model.Page
	Data ProductPage `json:"data"`
}

//ProductPage contains data specific to this page type
type ProductPage struct {
	HighlightedPages []HighlightedPage `json:"highlightedPages"`
	Children         []Child           `json:"children"`
}

//HighlightedPage contains some metadata for featured/highlight content
type HighlightedPage struct {
	Title string `json:"title"`
	URI   string `json:"uri"`
}

//Child contains metadata for child pages
type Child struct {
	Title       string `json:"title"`
	URI         string `json:"uri"`
	Description string `json:"description"`
}
