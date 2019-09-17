package render

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/render"
)

//Handler resolves the rendering of a specific pagem with a given model and template name
func Handler(w http.ResponseWriter, req *http.Request, page interface{}, page2 *model.Page, templateName string, f func()) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		render.JSON(w, 400, model.ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}

	req.Body.Close()

	err = json.Unmarshal(b, &page)
	if err != nil {
		render.JSON(w, 400, model.ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}

	log.Debug("incoming request json payload", log.Data{"json": string(b)})

	if f != nil {
		f()
	}

	page2.PatternLibraryAssetsPath = config.PatternLibraryAssetsPath

	page2.SiteDomain = config.SiteDomain

	log.DebugR(req, "rendered template", log.Data{"template": templateName})

	err = render.HTML(w, 200, templateName, page)
	if err != nil {
		render.JSON(w, 500, model.ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}
}
