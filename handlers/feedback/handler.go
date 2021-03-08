package feedback

import (
	"github.com/ONSdigital/dp-frontend-models/model/feedback"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"net/http"
)

func Handler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page feedback.Page

		render.Handler(w, req, &page, &page.Page, "feedback", nil, cfg)
	}
}
