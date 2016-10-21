package main

import (
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ONSdigital/dp-frontend-renderer/assets"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/homepage"
	"github.com/ONSdigital/dp-frontend-renderer/lang"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/ONSdigital/go-ns/handlers/healthcheck"
	"github.com/ONSdigital/go-ns/handlers/requestID"
	"github.com/ONSdigital/go-ns/handlers/timeout"
	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/pat"
	"github.com/justinas/alice"
	"github.com/nicksnyder/go-i18n/i18n"
	unrolled "github.com/unrolled/render"
)

func main() {
	bindAddr := os.Getenv("BIND_ADDR")
	if len(bindAddr) == 0 {
		bindAddr = ":20010"
	}

	err := lang.Load("en", "cy")
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
	log.Debug("languages", log.Data{"tags": i18n.LanguageTags()})

	enT, err := lang.Get("en")
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
	cyT, err := lang.Get("cy")
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
	log.Debug("translation", log.Data{"cy": cyT("You have {{.Count}} unread emails", 2), "en": enT("You have {{.Count}} unread emails", 2)})

	config.DebugMode, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Error(err, nil)
	}

	if config.DebugMode {
		config.PatternLibraryAssetsPath = "http://localhost:9000/dist"
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
