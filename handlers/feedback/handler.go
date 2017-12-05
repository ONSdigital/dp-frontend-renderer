package feedback

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/feedback"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	var page feedback.Page

	render.Handler(w, req, &page, &page.Page, "feedback", nil)
}
