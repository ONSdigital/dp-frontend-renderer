package datasetList

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dd/datasetList"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/ONSdigital/go-ns/log"
)

// Handler handles requests to render the list of all data discovery datasets. Expects the request to have a JSON body containing
// a marshalled datasetList.DatasetList model.
func Handler(w http.ResponseWriter, req *http.Request) {
	var page datasetList.DatasetList

	log.DebugR(req, "rendering datasetList page", nil)

	render.Handler(w, req, &page, &page.Page, "dd/datasetList", func() {
		log.DebugR(req, "datasetList page model", log.Data{"model": page})
	})
}
