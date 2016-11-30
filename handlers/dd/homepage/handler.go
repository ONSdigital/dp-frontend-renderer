package homepage

import (
	"net/http"
	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	var page model.Page

	render.Handler(w, req, &page, &page, "dd/homepage", nil)
}
