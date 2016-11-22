package main

import (
	"github.com/ONSdigital/go-ns/handlers/healthcheck"
	"github.com/ONSdigital/go-ns/handlers/requestID"
	"github.com/ONSdigital/go-ns/log"
	"github.com/carlhembrough/dp-api-spike/handlers/data"
	"github.com/carlhembrough/dp-api-spike/handlers/dataset"
	"github.com/carlhembrough/dp-api-spike/handlers/datasets"
	"github.com/carlhembrough/dp-api-spike/handlers/dimension"
	"github.com/carlhembrough/dp-api-spike/handlers/dimensions"
	"github.com/carlhembrough/dp-api-spike/handlers/download"
	"github.com/carlhembrough/dp-api-spike/handlers/search"
	"github.com/gorilla/pat"
	"github.com/justinas/alice"
	"net/http"
	"os"
)

func main() {
	bindAddr, found := os.LookupEnv("BIND_ADDR")
	if !found {
		bindAddr = ":20099"
	}

	log.Namespace = "dp-content-resolver"

	router := pat.New()
	alice := alice.New(log.Handler, requestID.Handler(16)).Then(router)

	router.Get("/healthcheck", healthcheck.Handler)

	router.Get("/datasets/search", search.Handler)
	router.Get("/datasets/{datasetId}/dimensions/{dimensionId}", dimension.Handler) // data
	router.Get("/datasets/{datasetId}/dimensions", dimensions.Handler)              // data
	router.Get("/datasets/{datasetId}/data", data.Handler)                          // data
	router.Get("/datasets/{datasetId}", dataset.Handler)                            // provide detailed information for a given dataset
	router.Get("/datasets", datasets.Handler)                                       // list high level dataset
	router.Get("/download", download.Handler)                                       // list high level dataset

	log.Debug("Starting server", log.Data{"bind_addr": bindAddr})

	if err := http.ListenAndServe(bindAddr, alice); err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
}
