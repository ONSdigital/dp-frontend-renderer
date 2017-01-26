package main

import (
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ONSdigital/dp-frontend-renderer/assets"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dd/dataset"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dd/datasetList"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dd/splash"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/homepage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/productPage"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/ONSdigital/go-ns/handlers/healthcheck"
	"github.com/ONSdigital/go-ns/handlers/requestID"
	"github.com/ONSdigital/go-ns/handlers/timeout"
	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/pat"
	"github.com/justinas/alice"
	unrolled "github.com/unrolled/render"
)

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
		config.DataDiscovery.AssetsPath = "http://localhost:20040"
	}

	log.Namespace = "dp-frontend-renderer"

	render.Renderer = unrolled.New(unrolled.Options{
		Asset:         assets.Asset,
		AssetNames:    assets.AssetNames,
		IsDevelopment: config.DebugMode,
		Layout:        "main",
		Funcs: []template.FuncMap{{
			"safeHTML": func(s string) template.HTML {
				return template.HTML(s)
			},
		}},
	})

	router := pat.New()
	alice := alice.New(
		timeout.Handler(10*time.Second),
		log.Handler,
		requestID.Handler(16),
	).Then(router)

	router.Get("/healthcheck", healthcheck.Handler)
	router.Post("/homepage", homepage.Handler)
	router.Post("/productPage", productPage.Handler)
	router.Post("/dd/datasetList", datasetList.Handler)
	router.Post("/dd/dataset", dataset.Handler)
	router.Post("/dd/splash", splash.Handler)

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
