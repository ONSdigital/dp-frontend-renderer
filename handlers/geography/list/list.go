package geographyList

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/geography/list"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page list.Page

	render.Handler(w, req, &page, &page.Page, "geography/list", nil)
}
