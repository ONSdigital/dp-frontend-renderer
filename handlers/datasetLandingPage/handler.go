package datasetLandingPage

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-renderer/config"

	"github.com/ONSdigital/dp-frontend-models/model/datasetVersionsList"

	"github.com/ONSdigital/dp-frontend-models/model/datasetEditionsList"
	"github.com/ONSdigital/dp-frontend-models/model/datasetLandingPageFilterable"
	"github.com/ONSdigital/dp-frontend-models/model/datasetLandingPageStatic"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

const xRequestIDParam = "X-Request-Id"

//DatasetHandler builds the template for dataset pages
func DatasetHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page datasetLandingPageStatic.Page

		render.Handler(w, req, &page, &page.Page, "dataset/dataset", nil, cfg)
	}
}

//FilterHandler ...
func FilterHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page datasetLandingPageFilterable.Page

		render.Handler(w, req, &page, &page.Page, "datasetLandingPage/filterable", nil, cfg)
	}
}

//StaticHandler builds the template for non-filterable dataset landing pages
func StaticHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page datasetLandingPageStatic.Page

		render.Handler(w, req, &page, &page.Page, "datasetLandingPage/static", nil, cfg)
	}
}

//NomisHandler ...
func NomisHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page datasetLandingPageFilterable.Page

		render.Handler(w, req, &page, &page.Page, "datasetLandingPage/nomis", nil, cfg)
	}
}

//EditionListHandler ...
func EditionListHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page datasetEditionsList.Page

		render.Handler(w, req, &page, &page.Page, "datasetLandingPage/edition-list", nil, cfg)
	}
}

func VersionListHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page datasetVersionsList.Page

		render.Handler(w, req, &page, &page.Page, "datasetLandingPage/version-list", nil, cfg)
	}
}
