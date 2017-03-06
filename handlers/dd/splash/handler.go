/*Package splash ...*/
package splash

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/dd/splash"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

//Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page splash.Page

	render.Handler(w, req, &page, &page.Page, "dd/splash", func() {
		page.HideHeaderAndFooter = true
	})
}
