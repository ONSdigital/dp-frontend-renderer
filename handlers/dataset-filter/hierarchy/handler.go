package hierarchy

import (
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/hierarchy"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler ...
func Handler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page hierarchy.Page

		render.Handler(w, req, &page, &page.Page, "dataset-filter/hierarchy", nil, cfg)
	}

}
