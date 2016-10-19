package model

type TaxonomyNode struct {
	Title    string         `json:"title"`
	URI      string         `json:"uri"`
	Children []TaxonomyNode `json:"children,omitempty"`
}
