package previewPage

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/previewPage"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page previewPage.Page

	render.Handler(w, req, &page, &page.Page, "dataset-filter/preview-page", nil)
}
