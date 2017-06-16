package middlePage

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/datasetpages/middlePage"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page middlePage.Page

	render.Handler(w, req, &page, &page.Page, "dataset/middlepage", nil)
}
