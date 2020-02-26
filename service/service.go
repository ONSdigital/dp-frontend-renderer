package service

import (
	"context"
	"github.com/ONSdigital/dp-frontend-renderer/assets"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	renderHelpers "github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/ONSdigital/dp-frontend-renderer/routes"
	"github.com/ONSdigital/go-ns/handlers/requestID"
	"github.com/ONSdigital/go-ns/handlers/timeout"
	deprecatedLog "github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/render"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/pat"
	"github.com/justinas/alice"
	"github.com/pkg/errors"
	unrolled "github.com/unrolled/render"
	"html/template"
	"net/http"
	"time"
)

const (
	namespace = "dp-frontend-renderer"
)

func Run(ctx context.Context, taxonomyRedirects map[string]string) error {
	log.Namespace = namespace

	cfg, err := config.Get()
	if err != nil {
		return errors.Wrap(err, "unable to retrieve service configuration")
	}

	log.Event(ctx, "got service configuration", log.Data{"config": cfg}, log.INFO)

	log.Event(ctx, "overriding default renderer with service assets", log.INFO)
	render.Renderer = unrolled.New(unrolled.Options{
		Asset:         assets.Asset,
		AssetNames:    assets.AssetNames,
		IsDevelopment: cfg.Debug,
		Layout:        "main",
		Funcs: []template.FuncMap{{
			"humanSize":                renderHelpers.HumanSize,
			"safeHTML":                 renderHelpers.SafeHTML,
			"dateFormat":               renderHelpers.DateFormat,
			"dateFormatYYYYMMDD":       renderHelpers.DateFormatYYYYMMDD,
			"last":                     renderHelpers.Last,
			"loop":                     renderHelpers.Loop,
			"subtract":                 renderHelpers.Subtract,
			"slug":                     renderHelpers.Slug,
			"markdown":                 renderHelpers.Markdown,
			"legacyDatasetDownloadURI": renderHelpers.LegacyDataSetDownloadURI,
			"localise":                 renderHelpers.Localise,
			"domainSetLang":            renderHelpers.DomainSetLang,
			"taxonomyLandingPage": func(s string) string {
				return taxonomyRedirects[s]
			},
		}},
	})

	router := pat.New()
	alice := alice.New(
		timeout.Handler(10*time.Second),
		deprecatedLog.Handler,
		requestID.Handler(16),
	).Then(router)
	routes.Setup(router, cfg)

	server := &http.Server{
		Addr:         cfg.BindAddr,
		Handler:      alice,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err = server.ListenAndServe(); err != nil {
		return errors.Wrap(err, "failed to start server")
	}
	// nil translates to exit code 0
	return nil
}
