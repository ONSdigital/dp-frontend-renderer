package ageSelectorList

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/ageSelectorList"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler handles rendering of age selector list template
func Handler(w http.ResponseWriter, req *http.Request) {
	var page ageSelectorList.Page

	render.Handler(w, req, &page, &page.Page, "dataset-filter/age-selector-list", nil)
}
