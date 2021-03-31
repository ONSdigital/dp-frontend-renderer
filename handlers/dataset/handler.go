package dataset

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-renderer/config"

	dataset "github.com/ONSdigital/dp-frontend-models/model/dataset"

	"github.com/ONSdigital/dp-frontend-renderer/render"
)

// Handler ...
func Handler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page dataset.Page

		render.Handler(w, req, &page, &page.Page, "dataset/test", nil, cfg)
	}
}
