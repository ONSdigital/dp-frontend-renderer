package ageSelector

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/age"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page age.Page

	render.Handler(w, req, &page, &page.Page, "dataset-filter/age", nil)
}
