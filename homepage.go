package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Homepage struct {
	Message string `json:"message"`
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
	if err != nil {
		r.JSON(w, 400, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if len(h.Message) == 0 {
		r.JSON(w, 400, ErrorResponse{
			Error: "'message' must be specified",
		})
		return
	}

	w.Header().Set("Content-Type", "text/html")
	r.HTML(w, 200, "homepage", h)
}
