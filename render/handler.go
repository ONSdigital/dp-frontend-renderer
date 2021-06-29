package render

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/go-ns/render"
	"github.com/ONSdigital/log.go/v2/log"
)

//Handler resolves the rendering of a specific pagem with a given model and template name
func Handler(w http.ResponseWriter, req *http.Request, page interface{}, page2 *model.Page, templateName string, f func(), cfg config.Config) {
	ctx := context.Background()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		render.JSON(w, 400, model.ErrorResponse{
			Error: err.Error(),
		})
		log.Error(ctx, "failed to read request body", err)
		return
	}

	req.Body.Close()

	err = json.Unmarshal(b, &page)
	if err != nil {
		render.JSON(w, 400, model.ErrorResponse{
			Error: err.Error(),
		})
		log.Error(ctx, "failed to unmarshal request body to page", err)
		return
	}

	if f != nil {
		f()
	}

	page2.PatternLibraryAssetsPath = cfg.PatternLibraryAssetsPath
	page2.SiteDomain = cfg.SiteDomain

	log.Info(ctx, "rendered template", log.Data{"template": templateName})

	err = render.HTML(w, 200, templateName, page)
	if err != nil {
		render.JSON(w, 500, model.ErrorResponse{
			Error: err.Error(),
		})
		log.Error(ctx, "failed to render template", err)
		return
	}
}
