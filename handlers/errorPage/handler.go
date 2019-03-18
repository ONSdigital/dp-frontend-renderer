/*Package errorPage ...*/
package errorPage

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/errorPage"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

const xRequestIDParam = "X-Request-Id"

//Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page errorPage.Page
	page.Metadata.ServiceName = "dp-frontend-router"
	render.Handler(w, req, &page, &page.Page, "error", nil)
}
