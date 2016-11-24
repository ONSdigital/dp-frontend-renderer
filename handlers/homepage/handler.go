/*Package homepage ...*/
package homepage

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-models/model/homepage"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

const xRequestIDParam = "X-Request-Id"

//Handler ...
func Handler(w http.ResponseWriter, req *http.Request) {
	var page homepage.Page

	render.Handler(w, req, &page, &page.Page, "homepage", func() {
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
