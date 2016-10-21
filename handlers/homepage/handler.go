/*Package homepage ...*/
package homepage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/lang"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/ONSdigital/go-ns/log"
)

//Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		render.JSON(w, 400, ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}

	req.Body.Close()

	var page Page
	err = json.Unmarshal(b, &page)
	if err != nil {
		render.JSON(w, 400, ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}

	if page.Data.HeadlineFigures != nil {
		for _, v := range page.Data.HeadlineFigures {
			if len(v.SparklineData) == 0 {
				continue
			}
			v.StartDate = v.SparklineData[0].Name
			v.EndDate = v.SparklineData[len(v.SparklineData)-1].Name
		}
	}

	page.Language = req.Header.Get("Accept-Language")
	if page.Language != "en" && page.Language != "cy" {
		page.Language = "en"
	}
	page.T, err = lang.Get(page.Language)
	if err != nil {
		render.JSON(w, 500, ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}

	page.PatternLibraryAssetsPath = config.PatternLibraryAssetsPath

	page.SiteDomain = config.SiteDomain

	log.DebugR(req, "rendered template", log.Data{"template": "homepage"})
	err = render.HTML(w, 200, "homepage", page)
	if err != nil {
		render.JSON(w, 500, ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}
}
