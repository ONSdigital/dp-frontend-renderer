// +build debug

package assets

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

// locales_active_cy_toml reads file data from disk. It returns an error on failure.
func locales_active_cy_toml() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/locales/active.cy.toml",
		"locales/active.cy.toml",
	)
}

// locales_active_en_toml reads file data from disk. It returns an error on failure.
func locales_active_en_toml() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/locales/active.en.toml",
		"locales/active.en.toml",
	)
}

// templates_cookies_preferences_tmpl reads file data from disk. It returns an error on failure.
func templates_cookies_preferences_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/cookies-preferences.tmpl",
		"templates/cookies-preferences.tmpl",
	)
}

// templates_dataset_filter_age_tmpl reads file data from disk. It returns an error on failure.
func templates_dataset_filter_age_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/dataset-filter/age.tmpl",
		"templates/dataset-filter/age.tmpl",
	)
}

// templates_dataset_filter_filter_overview_tmpl reads file data from disk. It returns an error on failure.
func templates_dataset_filter_filter_overview_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/dataset-filter/filter-overview.tmpl",
		"templates/dataset-filter/filter-overview.tmpl",
	)
}

// templates_dataset_filter_geography_tmpl reads file data from disk. It returns an error on failure.
func templates_dataset_filter_geography_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/dataset-filter/geography.tmpl",
		"templates/dataset-filter/geography.tmpl",
	)
}

// templates_dataset_filter_hierarchy_tmpl reads file data from disk. It returns an error on failure.
func templates_dataset_filter_hierarchy_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/dataset-filter/hierarchy.tmpl",
		"templates/dataset-filter/hierarchy.tmpl",
	)
}

// templates_dataset_filter_list_selector_tmpl reads file data from disk. It returns an error on failure.
func templates_dataset_filter_list_selector_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/dataset-filter/list-selector.tmpl",
		"templates/dataset-filter/list-selector.tmpl",
	)
}

// templates_dataset_filter_preview_page_tmpl reads file data from disk. It returns an error on failure.
func templates_dataset_filter_preview_page_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/dataset-filter/preview-page.tmpl",
		"templates/dataset-filter/preview-page.tmpl",
	)
}

// templates_dataset_filter_time_tmpl reads file data from disk. It returns an error on failure.
func templates_dataset_filter_time_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/dataset-filter/time.tmpl",
		"templates/dataset-filter/time.tmpl",
	)
}

// templates_datasetlandingpage_edition_list_tmpl reads file data from disk. It returns an error on failure.
func templates_datasetlandingpage_edition_list_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/datasetLandingPage/edition-list.tmpl",
		"templates/datasetLandingPage/edition-list.tmpl",
	)
}

// templates_datasetlandingpage_filterable_tmpl reads file data from disk. It returns an error on failure.
func templates_datasetlandingpage_filterable_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/datasetLandingPage/filterable.tmpl",
		"templates/datasetLandingPage/filterable.tmpl",
	)
}

// templates_datasetlandingpage_static_tmpl reads file data from disk. It returns an error on failure.
func templates_datasetlandingpage_static_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/datasetLandingPage/static.tmpl",
		"templates/datasetLandingPage/static.tmpl",
	)
}

// templates_datasetlandingpage_version_list_tmpl reads file data from disk. It returns an error on failure.
func templates_datasetlandingpage_version_list_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/datasetLandingPage/version-list.tmpl",
		"templates/datasetLandingPage/version-list.tmpl",
	)
}

// templates_error_tmpl reads file data from disk. It returns an error on failure.
func templates_error_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/error.tmpl",
		"templates/error.tmpl",
	)
}

// templates_feedback_tmpl reads file data from disk. It returns an error on failure.
func templates_feedback_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/feedback.tmpl",
		"templates/feedback.tmpl",
	)
}

// templates_geography_area_tmpl reads file data from disk. It returns an error on failure.
func templates_geography_area_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/geography/area.tmpl",
		"templates/geography/area.tmpl",
	)
}

// templates_geography_homepage_tmpl reads file data from disk. It returns an error on failure.
func templates_geography_homepage_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/geography/homepage.tmpl",
		"templates/geography/homepage.tmpl",
	)
}

// templates_geography_list_tmpl reads file data from disk. It returns an error on failure.
func templates_geography_list_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/geography/list.tmpl",
		"templates/geography/list.tmpl",
	)
}

// templates_homepage_tmpl reads file data from disk. It returns an error on failure.
func templates_homepage_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/homepage.tmpl",
		"templates/homepage.tmpl",
	)
}

// templates_main_tmpl reads file data from disk. It returns an error on failure.
func templates_main_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/main.tmpl",
		"templates/main.tmpl",
	)
}

// templates_partials_breadcrumb_tmpl reads file data from disk. It returns an error on failure.
func templates_partials_breadcrumb_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/partials/breadcrumb.tmpl",
		"templates/partials/breadcrumb.tmpl",
	)
}

// templates_partials_cookies_banner_tmpl reads file data from disk. It returns an error on failure.
func templates_partials_cookies_banner_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/partials/cookies-banner.tmpl",
		"templates/partials/cookies-banner.tmpl",
	)
}

// templates_partials_feedback_tmpl reads file data from disk. It returns an error on failure.
func templates_partials_feedback_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/partials/feedback.tmpl",
		"templates/partials/feedback.tmpl",
	)
}

// templates_partials_filter_selection_tmpl reads file data from disk. It returns an error on failure.
func templates_partials_filter_selection_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/partials/filter-selection.tmpl",
		"templates/partials/filter-selection.tmpl",
	)
}

// templates_partials_footer_tmpl reads file data from disk. It returns an error on failure.
func templates_partials_footer_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/partials/footer.tmpl",
		"templates/partials/footer.tmpl",
	)
}

// templates_partials_header_tmpl reads file data from disk. It returns an error on failure.
func templates_partials_header_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/partials/header.tmpl",
		"templates/partials/header.tmpl",
	)
}

// templates_partials_json_ld_base_tmpl reads file data from disk. It returns an error on failure.
func templates_partials_json_ld_base_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/partials/json-ld/base.tmpl",
		"templates/partials/json-ld/base.tmpl",
	)
}

// templates_partials_json_ld_dataset_common_tmpl reads file data from disk. It returns an error on failure.
func templates_partials_json_ld_dataset_common_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/partials/json-ld/dataset/common.tmpl",
		"templates/partials/json-ld/dataset/common.tmpl",
	)
}

// templates_partials_json_ld_dataset_filterable_tmpl reads file data from disk. It returns an error on failure.
func templates_partials_json_ld_dataset_filterable_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/partials/json-ld/dataset/filterable.tmpl",
		"templates/partials/json-ld/dataset/filterable.tmpl",
	)
}

// templates_partials_json_ld_dataset_legacy_tmpl reads file data from disk. It returns an error on failure.
func templates_partials_json_ld_dataset_legacy_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/partials/json-ld/dataset/legacy.tmpl",
		"templates/partials/json-ld/dataset/legacy.tmpl",
	)
}

// templates_partials_json_ld_dataset_timeseries_tmpl reads file data from disk. It returns an error on failure.
func templates_partials_json_ld_dataset_timeseries_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/partials/json-ld/dataset/timeseries.tmpl",
		"templates/partials/json-ld/dataset/timeseries.tmpl",
	)
}

// templates_partials_latest_release_alert_tmpl reads file data from disk. It returns an error on failure.
func templates_partials_latest_release_alert_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/partials/latest-release-alert.tmpl",
		"templates/partials/latest-release-alert.tmpl",
	)
}

// templates_partials_release_alert_tmpl reads file data from disk. It returns an error on failure.
func templates_partials_release_alert_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/partials/release-alert.tmpl",
		"templates/partials/release-alert.tmpl",
	)
}

// templates_partials_spinner_tmpl reads file data from disk. It returns an error on failure.
func templates_partials_spinner_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/partials/spinner.tmpl",
		"templates/partials/spinner.tmpl",
	)
}

// templates_productpage_tmpl reads file data from disk. It returns an error on failure.
func templates_productpage_tmpl() ([]byte, error) {
	return bindata_read(
		"/Users/david/gomod/src/github.com/ONSdigital/dp-frontend-renderer/assets/templates/productPage.tmpl",
		"templates/productPage.tmpl",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"locales/active.cy.toml": locales_active_cy_toml,
	"locales/active.en.toml": locales_active_en_toml,
	"templates/cookies-preferences.tmpl": templates_cookies_preferences_tmpl,
	"templates/dataset-filter/age.tmpl": templates_dataset_filter_age_tmpl,
	"templates/dataset-filter/filter-overview.tmpl": templates_dataset_filter_filter_overview_tmpl,
	"templates/dataset-filter/geography.tmpl": templates_dataset_filter_geography_tmpl,
	"templates/dataset-filter/hierarchy.tmpl": templates_dataset_filter_hierarchy_tmpl,
	"templates/dataset-filter/list-selector.tmpl": templates_dataset_filter_list_selector_tmpl,
	"templates/dataset-filter/preview-page.tmpl": templates_dataset_filter_preview_page_tmpl,
	"templates/dataset-filter/time.tmpl": templates_dataset_filter_time_tmpl,
	"templates/datasetLandingPage/edition-list.tmpl": templates_datasetlandingpage_edition_list_tmpl,
	"templates/datasetLandingPage/filterable.tmpl": templates_datasetlandingpage_filterable_tmpl,
	"templates/datasetLandingPage/static.tmpl": templates_datasetlandingpage_static_tmpl,
	"templates/datasetLandingPage/version-list.tmpl": templates_datasetlandingpage_version_list_tmpl,
	"templates/error.tmpl": templates_error_tmpl,
	"templates/feedback.tmpl": templates_feedback_tmpl,
	"templates/geography/area.tmpl": templates_geography_area_tmpl,
	"templates/geography/homepage.tmpl": templates_geography_homepage_tmpl,
	"templates/geography/list.tmpl": templates_geography_list_tmpl,
	"templates/homepage.tmpl": templates_homepage_tmpl,
	"templates/main.tmpl": templates_main_tmpl,
	"templates/partials/breadcrumb.tmpl": templates_partials_breadcrumb_tmpl,
	"templates/partials/cookies-banner.tmpl": templates_partials_cookies_banner_tmpl,
	"templates/partials/feedback.tmpl": templates_partials_feedback_tmpl,
	"templates/partials/filter-selection.tmpl": templates_partials_filter_selection_tmpl,
	"templates/partials/footer.tmpl": templates_partials_footer_tmpl,
	"templates/partials/header.tmpl": templates_partials_header_tmpl,
	"templates/partials/json-ld/base.tmpl": templates_partials_json_ld_base_tmpl,
	"templates/partials/json-ld/dataset/common.tmpl": templates_partials_json_ld_dataset_common_tmpl,
	"templates/partials/json-ld/dataset/filterable.tmpl": templates_partials_json_ld_dataset_filterable_tmpl,
	"templates/partials/json-ld/dataset/legacy.tmpl": templates_partials_json_ld_dataset_legacy_tmpl,
	"templates/partials/json-ld/dataset/timeseries.tmpl": templates_partials_json_ld_dataset_timeseries_tmpl,
	"templates/partials/latest-release-alert.tmpl": templates_partials_latest_release_alert_tmpl,
	"templates/partials/release-alert.tmpl": templates_partials_release_alert_tmpl,
	"templates/partials/spinner.tmpl": templates_partials_spinner_tmpl,
	"templates/productPage.tmpl": templates_productpage_tmpl,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"locales": &_bintree_t{nil, map[string]*_bintree_t{
		"active.cy.toml": &_bintree_t{locales_active_cy_toml, map[string]*_bintree_t{
		}},
		"active.en.toml": &_bintree_t{locales_active_en_toml, map[string]*_bintree_t{
		}},
	}},
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"cookies-preferences.tmpl": &_bintree_t{templates_cookies_preferences_tmpl, map[string]*_bintree_t{
		}},
		"dataset-filter": &_bintree_t{nil, map[string]*_bintree_t{
			"age.tmpl": &_bintree_t{templates_dataset_filter_age_tmpl, map[string]*_bintree_t{
			}},
			"filter-overview.tmpl": &_bintree_t{templates_dataset_filter_filter_overview_tmpl, map[string]*_bintree_t{
			}},
			"geography.tmpl": &_bintree_t{templates_dataset_filter_geography_tmpl, map[string]*_bintree_t{
			}},
			"hierarchy.tmpl": &_bintree_t{templates_dataset_filter_hierarchy_tmpl, map[string]*_bintree_t{
			}},
			"list-selector.tmpl": &_bintree_t{templates_dataset_filter_list_selector_tmpl, map[string]*_bintree_t{
			}},
			"preview-page.tmpl": &_bintree_t{templates_dataset_filter_preview_page_tmpl, map[string]*_bintree_t{
			}},
			"time.tmpl": &_bintree_t{templates_dataset_filter_time_tmpl, map[string]*_bintree_t{
			}},
		}},
		"datasetLandingPage": &_bintree_t{nil, map[string]*_bintree_t{
			"edition-list.tmpl": &_bintree_t{templates_datasetlandingpage_edition_list_tmpl, map[string]*_bintree_t{
			}},
			"filterable.tmpl": &_bintree_t{templates_datasetlandingpage_filterable_tmpl, map[string]*_bintree_t{
			}},
			"static.tmpl": &_bintree_t{templates_datasetlandingpage_static_tmpl, map[string]*_bintree_t{
			}},
			"version-list.tmpl": &_bintree_t{templates_datasetlandingpage_version_list_tmpl, map[string]*_bintree_t{
			}},
		}},
		"error.tmpl": &_bintree_t{templates_error_tmpl, map[string]*_bintree_t{
		}},
		"feedback.tmpl": &_bintree_t{templates_feedback_tmpl, map[string]*_bintree_t{
		}},
		"geography": &_bintree_t{nil, map[string]*_bintree_t{
			"area.tmpl": &_bintree_t{templates_geography_area_tmpl, map[string]*_bintree_t{
			}},
			"homepage.tmpl": &_bintree_t{templates_geography_homepage_tmpl, map[string]*_bintree_t{
			}},
			"list.tmpl": &_bintree_t{templates_geography_list_tmpl, map[string]*_bintree_t{
			}},
		}},
		"homepage.tmpl": &_bintree_t{templates_homepage_tmpl, map[string]*_bintree_t{
		}},
		"main.tmpl": &_bintree_t{templates_main_tmpl, map[string]*_bintree_t{
		}},
		"partials": &_bintree_t{nil, map[string]*_bintree_t{
			"breadcrumb.tmpl": &_bintree_t{templates_partials_breadcrumb_tmpl, map[string]*_bintree_t{
			}},
			"cookies-banner.tmpl": &_bintree_t{templates_partials_cookies_banner_tmpl, map[string]*_bintree_t{
			}},
			"feedback.tmpl": &_bintree_t{templates_partials_feedback_tmpl, map[string]*_bintree_t{
			}},
			"filter-selection.tmpl": &_bintree_t{templates_partials_filter_selection_tmpl, map[string]*_bintree_t{
			}},
			"footer.tmpl": &_bintree_t{templates_partials_footer_tmpl, map[string]*_bintree_t{
			}},
			"header.tmpl": &_bintree_t{templates_partials_header_tmpl, map[string]*_bintree_t{
			}},
			"json-ld": &_bintree_t{nil, map[string]*_bintree_t{
				"base.tmpl": &_bintree_t{templates_partials_json_ld_base_tmpl, map[string]*_bintree_t{
				}},
				"dataset": &_bintree_t{nil, map[string]*_bintree_t{
					"common.tmpl": &_bintree_t{templates_partials_json_ld_dataset_common_tmpl, map[string]*_bintree_t{
					}},
					"filterable.tmpl": &_bintree_t{templates_partials_json_ld_dataset_filterable_tmpl, map[string]*_bintree_t{
					}},
					"legacy.tmpl": &_bintree_t{templates_partials_json_ld_dataset_legacy_tmpl, map[string]*_bintree_t{
					}},
					"timeseries.tmpl": &_bintree_t{templates_partials_json_ld_dataset_timeseries_tmpl, map[string]*_bintree_t{
					}},
				}},
			}},
			"latest-release-alert.tmpl": &_bintree_t{templates_partials_latest_release_alert_tmpl, map[string]*_bintree_t{
			}},
			"release-alert.tmpl": &_bintree_t{templates_partials_release_alert_tmpl, map[string]*_bintree_t{
			}},
			"spinner.tmpl": &_bintree_t{templates_partials_spinner_tmpl, map[string]*_bintree_t{
			}},
		}},
		"productPage.tmpl": &_bintree_t{templates_productpage_tmpl, map[string]*_bintree_t{
		}},
	}},
}}
