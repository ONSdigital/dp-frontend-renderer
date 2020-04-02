package service

import (
	"context"
	"html/template"
	"net/http"
	"os"
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
	"github.com/pkg/errors"
	unrolled "github.com/unrolled/render"
)

const (
	namespace = "dp-frontend-renderer"
)

// Run creates the configuration, healthcheck and starts the server
func Run(ctx context.Context, taxonomyRedirects map[string]string, buildTime, gitCommit, version string) error {
	log.Namespace = namespace

	// Get Config
	cfg, err := config.Get()
	if err != nil {
		return errors.Wrap(err, "unable to retrieve service configuration")
	}
	log.Event(ctx, "got service configuration", log.Data{"config": cfg}, log.INFO)

	// Override default renderer with service assets
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

	// Create healthcheck with versionInfo
	versionInfo, err := healthcheck.NewVersionInfo(buildTime, gitCommit, version)
	if err != nil {
		log.Event(ctx, "failed to create versionInfo for healthcheck", log.FATAL, log.Error(err))
		os.Exit(1)
	}
	hc := healthcheck.New(versionInfo, cfg.HealthCheckRecoveryInterval, cfg.HealthCheckInterval)

	// Create router with middleware and routes
	router := pat.New()
	alice := alice.New(
		timeoutHandler(10*time.Second),
		log.Middleware,
		// deprecatedLog.Handler,
		requestID.Handler(16),
	).Then(router)
	routes.Setup(router, cfg, &hc)

	// Create HTTP server
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

// timeoutHandler implements a HTTP timeout
func timeoutHandler(timeout time.Duration) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.TimeoutHandler(h, timeout, "timed out")
	}
}
