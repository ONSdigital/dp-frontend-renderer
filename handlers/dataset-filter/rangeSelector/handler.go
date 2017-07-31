package rangeSelector

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/rangeSelector"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler handles rendering of range selector template
func Handler(w http.ResponseWriter, req *http.Request) {
	var page rangeSelector.Page

	render.Handler(w, req, &page, &page.Page, "dataset-filter/age-selector-range", nil)
}
