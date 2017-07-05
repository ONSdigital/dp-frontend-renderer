package geography

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/cmd/filterOverview"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page filterOverview.Page

	render.Handler(w, req, &page, &page.Page, "cmd/geography", nil)
}
