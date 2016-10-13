package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ONSdigital/dp-frontend-renderer/assets"
	"github.com/gorilla/pat"
	"github.com/unrolled/render"
)

var r *render.Render

func main() {
	bindAddr := os.Getenv("BIND_ADDR")
	if len(bindAddr) == 0 {
		bindAddr = ":8080"
	}

	r = render.New(render.Options{
		Asset:      assets.Asset,
		AssetNames: assets.AssetNames,
	})

	p := pat.New()

	p.Handle("/homepage", requestTime(homepage))

	if err := http.ListenAndServe(bindAddr, p); err != nil {
		log.Fatal(err)
	}
}

func requestTime(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		f(w, req)
		diff := time.Now().Sub(start)
		log.Printf("request completed: %v => %s", diff, req.URL.Path)
	}
}
