package productPage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/model"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/ONSdigital/go-ns/log"
)

//Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		render.JSON(w, 400, model.ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}

	req.Body.Close()

	var page Page
	err = json.Unmarshal(b, &page)
	if err != nil {
		render.JSON(w, 400, model.ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}

	page.Language = req.Header.Get("Accept-Language")
	if page.Language != "en" && page.Language != "cy" {
		page.Language = "en"
	}

	page.PatternLibraryAssetsPath = config.PatternLibraryAssetsPath

	page.SiteDomain = config.SiteDomain

	log.DebugR(req, "rendered template", log.Data{"template": "productPage"})
	err = render.HTML(w, 200, "productPage", page)
	if err != nil {
		render.JSON(w, 500, model.ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}
}
