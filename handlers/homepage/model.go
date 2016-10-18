package homepage

import model "github.com/onsdigital/dp-frontend-renderer/models"

type Request struct {
	Type       string               `json:"-"`
	URI        string               `json:"uri"`
	Taxonomy   []model.TaxonomyNode `json:"taxonomy"`
	Breadcrumb []model.TaxonomyNode `json:"breadcrumb"`
	Metadata   model.Metadata       `json:"metadata"`
	Data       Homepage             `json:"data"`
}

type Homepage struct {
	Publications    []Publication    `json:"publications"`
	HeadlineFigures []HeadlineFigure `json:"headlineFigures"`
	Featured        []Featured       `json:"featured"`
	Other           []Other          `json:"other"`
}

type Publication struct {
	Title       string `json:"title"`
	URI         string `json:"uri"`
	ReleaseDate string `json:"releaseDate"`
}

type HeadlineFigure struct {
	Title         string          `json:"title"`
	URI           string          `json:"uri"`
	ReleaseDate   string          `json:"releaseDate"`
	LatestFigure  LatestFigure    `json:"latestFigure"`
	SparklineData []SparklineData `json:"sparklineData"`
}

type SparklineData struct {
	Name    string  `json:"name"`
	Y       float32 `json:"y"`
	StringY string  `json:"stringY"`
}

type LatestFigure struct {
	PreUnit string `json:"preUnit"`
	Unit    string `json:"unit"`
	Figure  string `json:"figure"`
}

type Featured struct {
	Title string `json:"title"`
	URI   string `json:"uri"`
}

type Other struct {
	Title string `json:"title"`
	URI   string `json:"uri"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
