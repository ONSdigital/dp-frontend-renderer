package cookies

import (
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

//Handler ...
func Handler(cfg config.Config) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {
		var page model.Page

		render.Handler(w, req, &page, &page, "cookies-preferences", nil, cfg)
	}
}