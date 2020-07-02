package search

import (
	"github.com/ONSdigital/dp-frontend-models/model/search"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"net/http"
)

//Handler ...
func Handler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page search.Page

		render.Handler(w, req, &page, &page.Page, "search", nil, cfg)
	}
}
