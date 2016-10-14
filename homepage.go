package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Homepage struct {
	Type         string        `json:"type"`
	URI          string        `json:"uri"`
	Metadata     Metadata      `json:"metadata"`
	Publications []Publication `json:"publications"`
	Data         []Data        `json:"data"`
	Featured     []Featured    `json:"featured"`
	Other        []Other       `json:"other"`
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

func homepage(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		r.JSON(w, 400, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	req.Body.Close()

	var h Homepage
	err = json.Unmarshal(b, &h)
	fmt.Printf("%+v\n", &h)
	if err != nil {
		r.JSON(w, 400, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	// if len(h.Message) == 0 {
	// 	r.JSON(w, 400, ErrorResponse{
	// 		Error: "'message' must be specified",
	// 	})
	// 	return
	// }

	w.Header().Set("Content-Type", "text/html")
	r.HTML(w, 200, "homepage", h)
}
