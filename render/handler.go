package render

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/zebedee"
)

var ZebedeeClient *zebedee.Client

//Handler ... page and page2 need to be a pointer to the same value
func Handler(w http.ResponseWriter, req *http.Request, page interface{}, page2 *model.Page, templateName string, f func()) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		JSON(w, 400, model.ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}

	req.Body.Close()

	err = json.Unmarshal(b, &page)
	if err != nil {
		JSON(w, 400, model.ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}

	if f != nil {
		f()
	}

	page2.Language = req.Header.Get("Accept-Language")
	if page2.Language != "en" && page2.Language != "cy" {
		page2.Language = "en"
	}

	page2.PatternLibraryAssetsPath = config.PatternLibraryAssetsPath
	page2.SiteDomain = config.SiteDomain

	requestID := req.Header.Get("X-Request-Id")

	taxonomy, err := ZebedeeClient.GetTaxonomy(page2.URI, 2, requestID)
	if err != nil {
		JSON(w, 500, model.ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}

	for _, n := range taxonomy {
		page2.Taxonomy = append(page2.Taxonomy, n.Map())
	}

	parents, err := ZebedeeClient.GetParents(page2.URI, requestID)
	if err != nil {
		JSON(w, 500, model.ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}

	for _, n := range parents {
		page2.Breadcrumb = append(page2.Breadcrumb, n.Map())
	}

	log.DebugR(req, "rendered template", log.Data{"template": templateName})
	err = HTML(w, 200, templateName, page)
	if err != nil {
		JSON(w, 500, model.ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}
}
