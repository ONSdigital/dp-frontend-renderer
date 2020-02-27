package main

import (
	"context"
	"github.com/ONSdigital/dp-frontend-renderer/service"
	"github.com/ONSdigital/log.go/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var taxonomyRedirects map[string]string

func init() {
	// TODO: We should be building a breadcrumb from the API but for now
	// we will maintain a yaml file which redirects users to the taxonomy
	// base on the dataset ID
	if taxonomyRedirects == nil {
		taxonomyRedirects = make(map[string]string)
		b, err := ioutil.ReadFile("taxonomy-redirects.yml")
		if err != nil {
			log.Event(nil, "could not open taxonomy redirect file", log.Error(err), log.FATAL)
			os.Exit(1)
		}

		if err := yaml.Unmarshal(b, &taxonomyRedirects); err != nil {
			log.Event(nil, "could not unmarshal taxonomy redirects file", log.Error(err), log.FATAL)
			os.Exit(1)
		}
	}
}

func main() {
	ctx := context.Background()
	if err := service.Run(ctx, taxonomyRedirects); err != nil {
		log.Event(nil, "unable to run application", log.Error(err), log.FATAL)
		os.Exit(1)
	}
}
