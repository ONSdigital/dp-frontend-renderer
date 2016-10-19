/*Package homepage ...*/
package homepage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ONSdigital/dp-frontend-renderer/config"
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

	page.AssetsPath = config.AssetsPath

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
