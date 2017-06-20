package finishPage

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/datasetpages/finishPage"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page finishPage.Page

	render.Handler(w, req, &page, &page.Page, "dataset/finishpage", nil)
}
