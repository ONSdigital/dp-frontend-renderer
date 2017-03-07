/*Package disabled ...*/
package disabled

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dd/disabled"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

//Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page disabled.Page

	render.Handler(w, req, &page, &page.Page, "dd/disabled", nil)
}
