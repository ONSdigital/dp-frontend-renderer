package model

//Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	Type                     string         `json:"type"`
	URI                      string         `json:"uri"`
	Taxonomy                 []TaxonomyNode `json:"taxonomy"`
	Breadcrumb               []TaxonomyNode `json:"breadcrumb"`
	ServiceMessage           string         `json:"serviceMessage"`
	Metadata                 Metadata       `json:"metadata"`
	SiteDomain               string         `json:"-"`
	PatternLibraryAssetsPath string         `json:"-"`
	Language                 string         `json:"-"`
	HideHeaderAndFooter      bool         `json:"-"`
}
