package hello

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/hello"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

//Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	page := hello.Page{
		Greeting: "Hello, World!",
	}

	render.Handler(w, req, &page, &page.Page, "hello", nil)
}
