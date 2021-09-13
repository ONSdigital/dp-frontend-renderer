package routes

import (
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
	geographyArea "github.com/ONSdigital/dp-frontend-renderer/handlers/geography/area"
	geographyHomepage "github.com/ONSdigital/dp-frontend-renderer/handlers/geography/homepage"
	geographyList "github.com/ONSdigital/dp-frontend-renderer/handlers/geography/list"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/homepage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/productPage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/search"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/gorilla/pat"
)

func Setup(router *pat.Router, cfg *config.Config, hc *healthcheck.HealthCheck) {
	router.Get("/health", hc.Handler)
	router.Post("/homepage", homepage.Handler(*cfg))
	router.Post("/feedback", feedback.Handler(*cfg))
	router.Post("/dataset-landing-page-static", datasetLandingPage.StaticHandler(*cfg))
	router.Post("/dataset-edition-list", datasetLandingPage.EditionListHandler(*cfg))
	router.Post("/dataset-version-list", datasetLandingPage.VersionListHandler(*cfg))
	router.Post("/dataset-landing-page-filterable", datasetLandingPage.FilterHandler(*cfg))
	router.Post("/dataset-landing-page-nomis", datasetLandingPage.NomisHandler(*cfg))
	router.Post("/productPage", productPage.Handler(*cfg))
	router.Post("/error", errorPage.Handler(*cfg))
	router.Post("/dataset-filter/preview-page", previewPage.Handler(*cfg))
	router.Post("/dataset-filter/geography", geography.Handler(*cfg))
	router.Post("/dataset-filter/hierarchy", hierarchy.Handler(*cfg))
	router.Post("/dataset-filter/filter-overview", filterOverview.Handler(*cfg))
	router.Post("/dataset-filter/list-selector", listSelector.Handler(*cfg))
	router.Post("/dataset-filter/time", timeSelector.Handler(*cfg))
	router.Post("/dataset-filter/age", ageSelector.Handler(*cfg))
	router.Post("/geography-homepage", geographyHomepage.Handler(*cfg))
	router.Post("/geography-list", geographyList.Handler(*cfg))
	router.Post("/geography-area", geographyArea.Handler(*cfg))
	router.Post("/search", search.Handler(*cfg))
}
