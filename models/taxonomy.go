package model

type TaxonomyNode struct {
	Title    string         `json:"title"`
	Uri      string         `json:"uri"`
	Children []TaxonomyNode `json:"children,omitempty"`
}
