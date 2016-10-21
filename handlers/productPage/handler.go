package productPage

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-renderer/render"
)

//Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page Page

	render.Handler(w, req, &page, &page.Page, "productPage", nil)
}
