/*Package errorPage ...*/
package errorPage

import (
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/errorPage"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

const xRequestIDParam = "X-Request-Id"

//Handler ...
func Handler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var page errorPage.Page
		page.Metadata.ServiceName = "dp-frontend-renderer"
		render.Handler(w, req, &page, &page.Page, "error", nil, cfg)
	}

}
