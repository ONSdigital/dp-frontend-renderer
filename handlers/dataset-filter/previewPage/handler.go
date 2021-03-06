package previewPage

import (
	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/previewPage"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"net/http"
)

// Handler ...
func Handler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page previewPage.Page

		render.Handler(w, req, &page, &page.Page, "dataset-filter/preview-page", nil, cfg)
	}
}
