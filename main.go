package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/ONSdigital/dp-frontend-renderer/assets"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/ageSelector"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/filterOverview"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/geography"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/hierarchy"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/listSelector"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/previewPage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/timeSelector"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/datasetLandingPage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/errorPage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/feedback"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/geography/area"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/geography/homepage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/geography/list"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/homepage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/productPage"
	renderHelpers "github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/ONSdigital/go-ns/handlers/requestID"
	"github.com/ONSdigital/go-ns/handlers/timeout"
	"github.com/ONSdigital/go-ns/healthcheck"
	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/render"
	"github.com/gorilla/pat"
	"github.com/justinas/alice"
	unrolled "github.com/unrolled/render"
)

var taxonomyRedirects map[string]string

func init() {
	// TODO: We should be building a breadcrumb from the API but for now
	// we will maintain a yaml file which redirects users to the taxonomy
	// base on the dataset ID
	if taxonomyRedirects == nil {
		taxonomyRedirects = make(map[string]string)
		b, err := ioutil.ReadFile("taxonomy-redirects.yml")
		if err != nil {
			log.ErrorC("could not open taxonomy redirect file", err, nil)
			os.Exit(1)
		}

		if err := yaml.Unmarshal(b, &taxonomyRedirects); err != nil {
			log.ErrorC("could not unmarshal taxonomy redirects file", err, nil)
			os.Exit(1)
		}
	}
}

func main() {
	bindAddr := os.Getenv("BIND_ADDR")
	if len(bindAddr) == 0 {
		bindAddr = ":20010"
	}

	var err error
	config.DebugMode, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Error(err, nil)
	}

	if config.DebugMode {
		config.PatternLibraryAssetsPath = "http://localhost:9000/dist"
	}

	log.Namespace = "dp-frontend-renderer"

	log.Debug("overriding default renderer with service assets", nil)
	render.Renderer = unrolled.New(unrolled.Options{
		Asset:         assets.Asset,
		AssetNames:    assets.AssetNames,
		IsDevelopment: config.DebugMode,
		Layout:        "main",
		Funcs: []template.FuncMap{{
			"humanSize":          renderHelpers.HumanSize,
			"safeHTML":           renderHelpers.SafeHTML,
			"dateFormat":         renderHelpers.DateFormat,
			"dateFormatYYYYMMDD": renderHelpers.DateFormatYYYYMMDD,
			"last":               renderHelpers.Last,
			"loop":               renderHelpers.Loop,
			"subtract":           renderHelpers.Subtract,
			"slug":               renderHelpers.Slug,
			"legacyDatasetDownloadURI": renderHelpers.LegacyDataSetDownloadURI,
			"taxonomyLandingPage": func(s string) string {
				return taxonomyRedirects[s]
			},
		}},
	})

	router := pat.New()
	alice := alice.New(
		timeout.Handler(10*time.Second),
		log.Handler,
		requestID.Handler(16),
	).Then(router)

	router.Get("/healthcheck", healthcheck.Do)
	router.Post("/homepage", homepage.Handler)
	router.Post("/feedback", feedback.Handler)
	router.Post("/dataset-landing-page-static", datasetLandingPage.StaticHandler)
	router.Post("/dataset-edition-list", datasetLandingPage.EditionListHandler)
	router.Post("/dataset-version-list", datasetLandingPage.VersionListHandler)
	router.Post("/dataset-landing-page-filterable", datasetLandingPage.FilterHandler)
	router.Post("/productPage", productPage.Handler)
	router.Post("/error", errorPage.Handler)
	router.Post("/dataset-filter/preview-page", previewPage.Handler)
	router.Post("/dataset-filter/geography", geography.Handler)
	router.Post("/dataset-filter/hierarchy", hierarchy.Handler)
	router.Post("/dataset-filter/filter-overview", filterOverview.Handler)
	router.Post("/dataset-filter/list-selector", listSelector.Handler)
	router.Post("/dataset-filter/time", timeSelector.Handler)
	router.Post("/dataset-filter/age", ageSelector.Handler)
	router.Post("/geography-homepage", geographyHomepage.Handler)
	router.Post("/geography-list", geographyList.Handler)
	router.Post("/geography-area", geographyArea.Handler)

	log.Debug("Starting server", log.Data{"bind_addr": bindAddr})

	server := &http.Server{
		Addr:         bindAddr,
		Handler:      alice,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err = server.ListenAndServe(); err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
}
