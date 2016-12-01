package dataset

import (
	"github.com/ONSdigital/dp-frontend-models/model/dd/dataset"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"net/http"
)

// Handler handles requests to render a data discovery dataset.
// Expects the request body to contain the JSON version of a dataset.Page model object.
func Handler(w http.ResponseWriter, req *http.Request) {
	var page dataset.Page

	render.Handler(w, req, &page, &page.Page, "dd/dataset", nil)
}
