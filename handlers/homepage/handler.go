/*Package homepage ...*/
package homepage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/onsdigital/dp-frontend-renderer/render"
	"github.com/onsdigital/go-ns/log"
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

	var h Request
	err = json.Unmarshal(b, &h)
	if err != nil {
		render.JSON(w, 400, ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}

	log.DebugR(req, "rendered template", log.Data{"template": "homepage"})
	err = render.HTML(w, 200, "homepage", h)
	if err != nil {
		render.JSON(w, 500, ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}
}
