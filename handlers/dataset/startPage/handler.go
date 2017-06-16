package startPage

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/datasetpages/startPage"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page startPage.Page

	render.Handler(w, req, &page, &page.Page, "dataset/index", nil)
}
