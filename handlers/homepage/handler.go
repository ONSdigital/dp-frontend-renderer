/*Package homepage ...*/
package homepage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/model"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/ONSdigital/go-ns/log"
)

//Handler2 ...
func Handler2(w http.ResponseWriter, req *http.Request) {
	var page Page

	Handler(w, req, &page, &page.Page, "homepage", func() {
		if page.Data.HeadlineFigures != nil {
			for _, v := range page.Data.HeadlineFigures {
				if len(v.SparklineData) == 0 {
					continue
				}
				v.StartDate = v.SparklineData[0].Name
				v.EndDate = v.SparklineData[len(v.SparklineData)-1].Name
			}
		}
	})
}

//Handler ...
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

	f()

	page2.Language = req.Header.Get("Accept-Language")
	if page2.Language != "en" && page2.Language != "cy" {
		page2.Language = "en"
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
