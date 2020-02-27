package geographyHomepage

import (
	"github.com/ONSdigital/dp-frontend-models/model/geography/homepage"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"net/http"
)

// Handler ...
func Handler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page homepage.Page

		render.Handler(w, req, &page, &page.Page, "geography/homepage", nil, cfg)
	}
}
