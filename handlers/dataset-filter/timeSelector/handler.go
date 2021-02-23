package timeSelector

import (
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/time"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler ...
func Handler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page time.Page

		render.Handler(w, req, &page, &page.Page, "dataset-filter/time", nil, cfg)
	}

}
