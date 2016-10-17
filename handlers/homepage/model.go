package homepage

type Homepage struct {
	Type         string        `json:"type"`
	URI          string        `json:"uri"`
	Taxonomy     []Taxonomy    `json:"taxonomy"`
	Metadata     Metadata      `json:"metadata"`
	Publications []Publication `json:"publications"`
	Data         []Data        `json:"data"`
	Featured     []Featured    `json:"featured"`
	Other        []Other       `json:"other"`
}

type Taxonomy struct {
	Title    string  `json:"title"`
	Uri      string  `json:"uri"`
	Children []Child `json:"children"`
}

type Child struct {
	Title string `json:"title"`
	Uri   string `json:"uri"`
}

type Metadata struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    []string
}

type Publication struct {
	Title       string `json:"title"`
	URI         string `json:"uri"`
	ReleaseDate string `json:"releaseDate"`
}

type Data struct {
	Title          string         `json:"title"`
	URI            string         `json:"uri"`
	ReleaseDate    string         `json:"releaseDate"`
	HeadlineFigure HeadlineFigure `json:"headlineFigure"`
	Sparkline      []Sparkline    `json:"sparklineData"`
}

type Sparkline struct {
	Name    string  `json:"name"`
	Y       float32 `json:"y"`
	StringY string  `json:"stringY"`
}

type HeadlineFigure struct {
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
