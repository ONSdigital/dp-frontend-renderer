package errorPage

import "github.com/ONSdigital/dp-frontend-models/model"

//Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	model.Page
	Error Error `json:"error"`
}

type Error struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
