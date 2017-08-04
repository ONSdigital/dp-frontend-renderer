package listSelector

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/listSelector"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler handles rendering of age selector list template
func Handler(w http.ResponseWriter, req *http.Request) {
	var page listSelector.Page

	render.Handler(w, req, &page, &page.Page, "dataset-filter/list-selector", nil)
}
