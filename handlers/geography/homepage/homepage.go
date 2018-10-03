package geographyHomepage

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/geography/homepage"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page homepage.Page

	render.Handler(w, req, &page, &page.Page, "geography/homepage", nil)
}
