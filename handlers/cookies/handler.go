package cookies

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-renderer/config"

	"github.com/ONSdigital/dp-frontend-models/model/cookiespreferences"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

//Handler ...
func Handler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page cookiespreferences.Page

		render.Handler(w, req, &page, &page.Page, "cookies-preferences", nil, cfg)
	}
}
