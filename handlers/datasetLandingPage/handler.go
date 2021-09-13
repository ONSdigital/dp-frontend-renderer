package datasetLandingPage

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-renderer/config"

	"github.com/ONSdigital/dp-frontend-models/model/datasetVersionsList"

	"github.com/ONSdigital/dp-frontend-models/model/datasetLandingPageStatic"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

const xRequestIDParam = "X-Request-Id"

//StaticHandler builds the template for non-filterable dataset landing pages
func StaticHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page datasetLandingPageStatic.Page

		render.Handler(w, req, &page, &page.Page, "datasetLandingPage/static", nil, cfg)
	}
}

func VersionListHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page datasetVersionsList.Page

		render.Handler(w, req, &page, &page.Page, "datasetLandingPage/version-list", nil, cfg)
	}
}
