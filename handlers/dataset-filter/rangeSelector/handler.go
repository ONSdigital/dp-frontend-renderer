package rangeSelector

import (
	"net/http"
	"time"

	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/rangeSelector"
	"github.com/ONSdigital/dp-frontend-renderer/render"
)

type TimeSlice []time.Time

func (p TimeSlice) Len() int {
	return len(p)
}

func (p TimeSlice) Less(i, j int) bool {
	return p[i].Before(p[j])
}

func (p TimeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Handler handles rendering of range selector template
func Handler(w http.ResponseWriter, req *http.Request) {
	var page rangeSelector.Page

	render.Handler(w, req, &page, &page.Page, "dataset-filter/range-selector", nil)
}
