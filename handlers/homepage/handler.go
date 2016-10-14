package homepage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/ONSdigital/go-ns/log"
)

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

	var h Homepage
	err = json.Unmarshal(b, &h)
	fmt.Printf("%+v\n", &h)
	if err != nil {
		render.JSON(w, 400, ErrorResponse{
			Error: err.Error(),
		})
		log.ErrorR(req, err, nil)
		return
	}

	// if len(h.Message) == 0 {
	// 	r.JSON(w, 400, ErrorResponse{
	// 		Error: "'message' must be specified",
	// 	})
	// 	return
	// }

	log.DebugR(req, "rendered template", log.Data{"template": "homepage"})
	w.Header().Set("Content-Type", "text/html")
	render.HTML(w, 200, "homepage", h)
}
