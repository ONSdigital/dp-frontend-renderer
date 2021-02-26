package geographyArea

import (
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/geography/area"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler ...
func Handler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page area.Page

		render.Handler(w, req, &page, &page.Page, "geography/area", nil, cfg)
	}
}
