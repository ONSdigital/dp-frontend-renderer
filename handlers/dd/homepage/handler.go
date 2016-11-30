package homepage

import (
	"net/http"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/ONSdigital/dp-frontend-models/model/dd/homepage"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	var page homepage.Homepage

	render.Handler(w, req, &page, &page.Page, "dd/homepage", nil)
}
