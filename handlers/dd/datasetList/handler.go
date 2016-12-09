package datasetList

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dd/datasetList"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler handles requests to render the list of all data discovery datasets. Expects the request to have a JSON body containing
// a marshalled datasetList.DatasetList model.
func Handler(w http.ResponseWriter, req *http.Request) {
	var page datasetList.DatasetList

	render.Handler(w, req, &page, &page.Page, "dd/datasetList", nil)
}
