package ageSelectorRange

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/ageSelectorRange"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler handles rendering of age selector range template
func Handler(w http.ResponseWriter, req *http.Request) {
	var page ageSelectorRange.Page

	render.Handler(w, req, &page, &page.Page, "dataset-filter/age-selector-range", nil)
}
