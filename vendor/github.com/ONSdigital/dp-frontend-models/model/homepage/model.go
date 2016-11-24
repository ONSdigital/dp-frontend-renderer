package homepage

import "github.com/ONSdigital/dp-frontend-models/model"

//Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	model.Page
	Data Homepage `json:"data"`
}

//Homepage contains data specific to this page type
type Homepage struct {
	Releases        []Release         `json:"releases"`
	HeadlineFigures []*HeadlineFigure `json:"headlineFigures"`
	Featured        []Featured        `json:"featured"`
	Other           []Other           `json:"other"`
}

//Release is the data for an individual release
type Release struct {
	Title       string `json:"title"`
	URI         string `json:"uri"`
	ReleaseDate string `json:"releaseDate"`
}

//HeadlineFigure is the data for an individual timeseries
type HeadlineFigure struct {
	Title         string          `json:"title"`
	URI           string          `json:"uri"`
	ReleaseDate   string          `json:"releaseDate"`
	LatestFigure  LatestFigure    `json:"latestFigure"`
	SparklineData []SparklineData `json:"sparklineData"`
	StartDate     string          `json:"-"`
	EndDate       string          `json:"-"`
}

//SparklineData is the data that is sent to highcharts to produce the sparkline for each timeseries
type SparklineData struct {
	Name    string  `json:"name"`
	Y       float32 `json:"y"`
	StringY string  `json:"stringY"`
}

//LatestFigure is the extra information displayed for the latest figure for a timeseries
type LatestFigure struct {
	PreUnit string `json:"preUnit"`
	Unit    string `json:"unit"`
	Figure  string `json:"figure"`
}

//Featured is data for content that displays under a 'featured' heading
type Featured struct {
	Title string `json:"title"`
	URI   string `json:"uri"`
}

//Other is data for content that displays under a 'other' heading
type Other struct {
	Title string `json:"title"`
	URI   string `json:"uri"`
}
