package geographyHomepage

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/geographyHomepage"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// HomepageHandler ...
func HomepageHandler(w http.ResponseWriter, req *http.Request) {
	var page geographyHomepage.Page

	render.Handler(w, req, &page, &page.Page, "geography-homepage", nil)
}
