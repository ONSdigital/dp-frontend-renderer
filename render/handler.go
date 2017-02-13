package render

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/zebedee"
	zebedeeModel "github.com/ONSdigital/go-ns/zebedee/model"
)

var ZebedeeClient *zebedee.Client
var taxonomyCache *[]model.TaxonomyNode
var serviceMessage *string
var taxonomyMutex sync.Mutex
var serviceMessageMutex sync.Mutex

func getServiceMessage(requestID string) (string, error) {
	if serviceMessage != nil {
		return *serviceMessage, nil
	}

	serviceMessageMutex.Lock()
	defer serviceMessageMutex.Unlock()

	if serviceMessage != nil {
		return *serviceMessage, nil
	}

	data, _, err := ZebedeeClient.GetData("/", requestID)
	if err != nil {
		return "", err
	}

	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return "", err
	}

	sM := m["serviceMessage"]
	if s, ok := sM.(string); ok {
		serviceMessage = &s
	}

	if serviceMessage != nil {
		return *serviceMessage, nil
	}

	return "", nil
}

func getTaxonomy(requestID string) ([]model.TaxonomyNode, error) {
	if taxonomyCache != nil {
		return *taxonomyCache, nil
	}

	taxonomyMutex.Lock()
	defer taxonomyMutex.Unlock()

	if taxonomyCache != nil {
		return *taxonomyCache, nil
	}

	zebedeeTaxonomy, err := ZebedeeClient.GetTaxonomy("/", 2, requestID)
	if err != nil {
		return nil, err
	}

	var taxonomy []model.TaxonomyNode
	for _, n := range zebedeeTaxonomy {
		if n.PageType == zebedee.TaxonomyLandingPage {
			taxonomy = append(taxonomy, n.Map())
		}
	}

	taxonomyCache = &taxonomy
	return taxonomy, nil
}

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

	wg := new(sync.WaitGroup)

	wg.Add(3)

	go func() {
		defer wg.Done()

		var serviceMessage string
		serviceMessage, err = getServiceMessage(requestID)
		if err != nil {
			log.ErrorR(req, err, nil)
		}

		page2.ServiceMessage = serviceMessage
	}()

	go func() {
		defer wg.Done()

		var taxonomy []model.TaxonomyNode
		taxonomy, err = getTaxonomy(requestID)
		if err != nil {
			log.ErrorR(req, err, nil)
		}
		page2.Taxonomy = taxonomy
	}()

	go func() {
		defer wg.Done()

		var parents []zebedeeModel.ContentNode
		parents, err = ZebedeeClient.GetParents(req.URL.Path, requestID)
		if err != nil {
			log.ErrorR(req, err, nil)
		}

		for _, n := range parents {
			page2.Breadcrumb = append(page2.Breadcrumb, n.Map())
		}
	}()

	wg.Wait()

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
