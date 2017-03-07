package dataset

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dd"
	"github.com/ONSdigital/dp-frontend-models/model/dd/dataset"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler handles requests to render a data discovery dataset.
// Expects the request body to contain the JSON version of a dataset.Page model object.
func Handler(w http.ResponseWriter, req *http.Request) {
	var page dataset.Page

	render.Handler(w, req, &page, &page.Page, "dd/dataset", func() {
		// TODO Add in graceful error handling in this callback function
		page.Dataset.Config = &dd.Config{
			ReactAppAssetsPath: config.DataDiscovery.ReactAppAssetsPath,
			ApiURL:           config.DataDiscovery.ApiURL,
			JobsApiURL:       config.DataDiscovery.JobsApiURL,
			BasePath:         config.DataDiscovery.BasePath,
		}
	})
}
