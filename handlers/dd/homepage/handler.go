package homepage

import (
	"github.com/ONSdigital/dp-frontend-models/model/dd/homepage"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	var page homepage.Homepage

	render.Handler(w, req, &page, &page.Page, "dd/homepage", nil)
}
