package geographyArea

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/geography/area"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page area.Page

	render.Handler(w, req, &page, &page.Page, "geography/area", nil)
}
