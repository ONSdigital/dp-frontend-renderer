package productPage

import "github.com/ONSdigital/dp-frontend-renderer/model"

//Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	model.Page
	Data ProductPage `json:"data"`
}

//ProductPage contains data specific to this page type
type ProductPage struct {
}
