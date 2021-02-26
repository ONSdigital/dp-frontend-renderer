package service

import (
	"context"
	"html/template"
	"net/http"
	"time"

	"github.com/ONSdigital/dp-frontend-renderer/assets"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	renderHelpers "github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/ONSdigital/dp-frontend-renderer/routes"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/go-ns/handlers/requestID"
	"github.com/ONSdigital/go-ns/render"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/pat"
	"github.com/justinas/alice"
	unrolled "github.com/unrolled/render"
)

// Run creates the routes, middleware and starts the HTTP server in a separate go routine.
func Run(ctx context.Context, cfg *config.Config, taxonomyRedirects map[string]string, hc *healthcheck.HealthCheck) *http.Server {

	// Override default renderer with service assets
	log.Event(ctx, "overriding default renderer with service assets", log.INFO)
	render.Renderer = unrolled.New(unrolled.Options{
		Asset:         assets.Asset,
		AssetNames:    assets.AssetNames,
		IsDevelopment: cfg.Debug,
		Layout:        "main",
		Funcs: []template.FuncMap{{
			"humanSize":                   renderHelpers.HumanSize,
			"safeHTML":                    renderHelpers.SafeHTML,
			"dateFormat":                  renderHelpers.DateFormat,
			"dateFormatYYYYMMDD":          renderHelpers.DateFormatYYYYMMDD,
			"DatePeriodFormat":            renderHelpers.DatePeriodFormat,
			"last":                        renderHelpers.Last,
			"loop":                        renderHelpers.Loop,
			"add":                         renderHelpers.Add,
			"subtract":                    renderHelpers.Subtract,
			"slug":                        renderHelpers.Slug,
			"markdown":                    renderHelpers.Markdown,
			"legacyDatasetDownloadURI":    renderHelpers.LegacyDataSetDownloadURI,
			"localise":                    renderHelpers.Localise,
			"domainSetLang":               renderHelpers.DomainSetLang,
			"hasField":                    renderHelpers.HasField,
			"concatenateStrings":          renderHelpers.ConcatenateStrings,
			"notLastItem":                 renderHelpers.NotLastItem,
			"truncateToMaximumCharacters": renderHelpers.TruncateToMaximumCharacters,
			"taxonomyLandingPage": func(s string) string {
				return taxonomyRedirects[s]
			},
		}},
	})

	// Create router with middleware and routes
	router := pat.New()
	alice := alice.New(
		timeoutHandler(10*time.Second),
		log.Middleware,
		requestID.Handler(16),
	).Then(router)
	routes.Setup(router, cfg, hc)

	// Create HTTP server
	server := &http.Server{
		Addr:         cfg.BindAddr,
		Handler:      alice,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start the HTTP server in a new go routine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Event(ctx, "error starting http server", log.ERROR, log.Error(err))
		}
	}()

	return server
}

// timeoutHandler implements a HTTP timeout
func timeoutHandler(timeout time.Duration) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.TimeoutHandler(h, timeout, "timed out")
	}
}
