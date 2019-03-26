package filterOverview

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/filterOverview"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler handles rendering of filter overview template
func Handler(w http.ResponseWriter, req *http.Request) {
	var page filterOverview.Page

	render.Handler(w, req, &page, &page.Page, "dataset-filter/filter-overview", nil)
}
