package cookies

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

//Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page model.Page

	render.Handler(w, req, &page, &page, "cookies-preferences", nil)

}
