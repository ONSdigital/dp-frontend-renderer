package dd

// Datasets represents a set of available datasets.
type Datasets struct {
	Items        []*Dataset `json:"items,omitempty"`
	Count        int        `json:"count"`
	Total        int        `json:"total"`
	StartIndex   int        `json:"startIndex"`
	ItemsPerPage int        `json:"itemsPerPage"`
}

// Dataset represents the details about a particular dataset.
type Dataset struct {
	DataDiscoveryAssetsPath string       `json:"-"` // Path to get React bundle.js from - only needed on dataset pages
	ID                      string       `json:"id"`
	Title                   string       `json:"title"`
	URL                     string       `json:"url,omitempty"`
	Metadata                *Metadata    `json:"metadata,omitempty"`
	Dimensions              []*Dimension `json:"dimensions,omitempty"`
	Data                    *Table       `json:"data,omitempty"`
}

// Metadata about a dataset.
type Metadata struct {
	Description string   `json:"description,omitempty"`
	Taxonomies  []string `json:"taxonomies,omitempty"`
}

// Dimension of a dataset.
type Dimension struct {
	ID   string `json:"id"`
	Name string `json:"name"` // Sex

	Options        []*DimensionOption `json:"options,omitempty"`
	SelectedOption *DimensionOption   `json:"selectedOption,omitempty"`
}

// DimensionOption describes a possible value for a dimension.
type DimensionOption struct {
	ID   string `json:"id"`
	Name string `json:"name"` // Male

	Options []*DimensionOption `json:"options,omitempty"`
}

// Row is a particular data point in a dataset.
type Row struct {
	Observation interface{}        // 123
	Dimensions  []*DimensionOption // Sex=Male
}

// Table is a set of rows in a dataset.
type Table struct {
	Rows []*Row
}
