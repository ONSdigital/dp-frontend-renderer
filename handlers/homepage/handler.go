/*Package homepage ...*/
package homepage

import (
	"github.com/ONSdigital/dp-frontend-models/model/geography/homepage"
	"net/http"

	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

const xRequestIDParam = "X-Request-Id"

//Handler ...
func Handler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page homepage.Page

		render.Handler(w, req, &page, &page.Page, "homepage", nil, cfg)

	}
}
