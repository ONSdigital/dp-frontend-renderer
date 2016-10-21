package productPage

import "github.com/ONSdigital/dp-frontend-renderer/model"

//Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	Type                     string               `json:"type"`
	URI                      string               `json:"uri"`
	Taxonomy                 []model.TaxonomyNode `json:"taxonomy"`
	Breadcrumb               []model.TaxonomyNode `json:"breadcrumb"`
	ServiceMessage           string               `json:"serviceMessage"`
	Metadata                 model.Metadata       `json:"metadata"`
	SiteDomain               string               `json:"-"`
	PatternLibraryAssetsPath string               `json:"-"`
	Language                 string               `json:"-"`
	Data                     ProductPage          `json:"data"`
}

//ProductPage contains data specific to this page type
type ProductPage struct {
}
