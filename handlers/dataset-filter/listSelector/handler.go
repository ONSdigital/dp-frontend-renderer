package listSelector

import (
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/listSelector"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler handles rendering of age selector list template
func Handler(cfg config.Config) http.HandlerFunc{
	return func (w http.ResponseWriter, req *http.Request) {
		var page listSelector.Page

		render.Handler(w, req, &page, &page.Page, "dataset-filter/list-selector", nil, cfg)
	}

}
