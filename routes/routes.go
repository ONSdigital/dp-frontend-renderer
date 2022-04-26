package routes

import (
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/cookies"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/ageSelector"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/filterOverview"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/geography"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/hierarchy"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/listSelector"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/previewPage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/timeSelector"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/errorPage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/feedback"
	geographyArea "github.com/ONSdigital/dp-frontend-renderer/handlers/geography/area"
	geographyHomepage "github.com/ONSdigital/dp-frontend-renderer/handlers/geography/homepage"
	geographyList "github.com/ONSdigital/dp-frontend-renderer/handlers/geography/list"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/homepage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/productPage"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/gorilla/mux"
	"net/http"
)

func Setup(router *mux.Router, cfg *config.Config, hc *healthcheck.HealthCheck) {
	router.Path("/health").HandlerFunc(hc.Handler).Methods(http.MethodGet)
	router.Path("/homepage").Handler(homepage.Handler(*cfg)).Methods(http.MethodPost)
	router.Path("/homepage").Handler(homepage.Handler(*cfg)).Methods(http.MethodPost)
	router.Path("/feedback").Handler(feedback.Handler(*cfg)).Methods(http.MethodPost)
	router.Path("/productPage").Handler(productPage.Handler(*cfg)).Methods(http.MethodPost)
	router.Path("/error").Handler(errorPage.Handler(*cfg)).Methods(http.MethodPost)
	router.Path("/dataset-filter/preview-page").Handler(previewPage.Handler(*cfg)).Methods(http.MethodPost)
	router.Path("/dataset-filter/geography").Handler(geography.Handler(*cfg)).Methods(http.MethodPost)
	router.Path("/dataset-filter/hierarchy").Handler(hierarchy.Handler(*cfg)).Methods(http.MethodPost)
	router.Path("/dataset-filter/filter-overview").Handler(filterOverview.Handler(*cfg)).Methods(http.MethodPost)
	router.Path("/dataset-filter/list-selector").Handler(listSelector.Handler(*cfg)).Methods(http.MethodPost)
	router.Path("/dataset-filter/time").Handler(timeSelector.Handler(*cfg)).Methods(http.MethodPost)
	router.Path("/dataset-filter/age").Handler(ageSelector.Handler(*cfg)).Methods(http.MethodPost)
	router.Path("/geography-homepage").Handler(geographyHomepage.Handler(*cfg)).Methods(http.MethodPost)
	router.Path("/geography-list").Handler(geographyList.Handler(*cfg)).Methods(http.MethodPost)
	router.Path("/geography-area").Handler(geographyArea.Handler(*cfg)).Methods(http.MethodPost)
	router.Path("/cookies-preferences").Handler(cookies.Handler(*cfg)).Methods(http.MethodPost)
}
