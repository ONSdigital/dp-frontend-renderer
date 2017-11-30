package datasetLandingPage

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/datasetVersionsList"

	"github.com/ONSdigital/dp-frontend-models/model/datasetEditionsList"
	"github.com/ONSdigital/dp-frontend-models/model/datasetLandingPageFilterable"
	"github.com/ONSdigital/dp-frontend-models/model/datasetLandingPageStatic"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

const xRequestIDParam = "X-Request-Id"

//FilterHandler ...
func FilterHandler(w http.ResponseWriter, req *http.Request) {
	var page datasetLandingPageFilterable.Page

	render.Handler(w, req, &page, &page.Page, "datasetLandingPage/filterable", nil)
}

//StaticHandler builds the template for non-filterable dataset landing pages
func StaticHandler(w http.ResponseWriter, req *http.Request) {
	var page datasetLandingPageStatic.Page

	render.Handler(w, req, &page, &page.Page, "datasetLandingPage/static", nil)
}

//EditionListHandler ...
func EditionListHandler(w http.ResponseWriter, req *http.Request) {
	var page datasetEditionsList.Page

	render.Handler(w, req, &page, &page.Page, "datasetLandingPage/edition-list", nil)
}

func VersionListHandler(w http.ResponseWriter, req *http.Request) {
	var page datasetVersionsList.Page

	render.Handler(w, req, &page, &page.Page, "datasetLandingPage/version-list", nil)
}
