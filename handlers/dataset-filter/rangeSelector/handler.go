package rangeSelector

import (
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"time"

	"github.com/ONSdigital/dp-frontend-models/model/dataset-filter/rangeSelector"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/murlokswarm/log"
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

	render.Handler(w, req, &page, &page.Page, "dataset-filter/range-selector", func() {
		var dates []time.Time
		for _, val := range page.Data.RangeData.Values {
			myrReg := regexp.MustCompile(`^(\d{4})\.(\d{1}|\d{2})$`)
			myrSubs := myrReg.FindStringSubmatch(val)
			if len(myrSubs) == 3 {
				date, err := time.Parse("01-02-2006", fmt.Sprintf("%02s-01-%s", myrSubs[2], myrSubs[1]))
				if err != nil {
					log.Error(err, nil)
				}
				dates = append(dates, date)
			}
		}
		if len(dates) > 0 {
			sortedDates := TimeSlice(dates)
			sort.Sort(sortedDates)
			for i, d := range sortedDates {
				page.Data.RangeData.Values[i] = fmt.Sprintf("%s %d", d.Month().String(), d.Year())
			}
		}
	})
}
