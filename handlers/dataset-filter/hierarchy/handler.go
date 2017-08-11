package hierarchy

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/hierarchy"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page hierarchy.Page

	render.Handler(w, req, &page, &page.Page, "dataset-filter/hierarchy", nil)
}
