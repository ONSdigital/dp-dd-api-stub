package main

import (
	"net/http"
	"os"

	"github.com/ONSdigital/dp-dd-api-stub/handlers/data"
	"github.com/ONSdigital/dp-dd-api-stub/handlers/dataresource"
	"github.com/ONSdigital/dp-dd-api-stub/handlers/dataset"
	"github.com/ONSdigital/dp-dd-api-stub/handlers/datasets"
	"github.com/ONSdigital/dp-dd-api-stub/handlers/dimension"
	"github.com/ONSdigital/dp-dd-api-stub/handlers/dimensions"
	"github.com/ONSdigital/dp-dd-api-stub/handlers/download"
	"github.com/ONSdigital/dp-dd-api-stub/handlers/geoarea"
	"github.com/ONSdigital/dp-dd-api-stub/handlers/geoareas"
	"github.com/ONSdigital/dp-dd-api-stub/handlers/geoareatype"
	"github.com/ONSdigital/dp-dd-api-stub/handlers/geohierarchies"
	"github.com/ONSdigital/dp-dd-api-stub/handlers/geohierarchy"
	"github.com/ONSdigital/dp-dd-api-stub/handlers/hierarchies"
	"github.com/ONSdigital/dp-dd-api-stub/handlers/hierarchy"
	"github.com/ONSdigital/dp-dd-api-stub/handlers/search"
	"github.com/ONSdigital/go-ns/handlers/healthcheck"
	"github.com/ONSdigital/go-ns/handlers/requestID"
	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/pat"
	"github.com/justinas/alice"
)

func main() {
	bindAddr, found := os.LookupEnv("BIND_ADDR")
	if !found {
		bindAddr = ":20099"
	}

	log.Namespace = "dp-dd-api-stub"

	router := pat.New()
	alice := alice.New(log.Handler, requestID.Handler(16), corsHandler).Then(router)

	router.Get("/healthcheck", healthcheck.Handler)

	// Datasets
	router.Options("/{x:.*}", func(w http.ResponseWriter, r *http.Request) {})

	router.Get("/datasets/search", search.Handler)
	// datasets/{datasetId}/dimensions/{dimensionId}			<- dimension with list of options
	router.Get("/datasets/{datasetId}/dimensions/{dimensionId}", dimension.Handler) // data

	router.Get("/datasets/{datasetId}/dimensions", dimensions.Handler) // data
	router.Get("/datasets/{datasetId}/data", data.Handler)             // data
	router.Get("/datasets/{datasetId}", dataset.Handler)               // provide detailed information for a given dataset
	router.Get("/datasets", datasets.Handler)                          // list high level dataset

	router.Get("/hierarchies/{hierarchyId}", hierarchy.Handler)
	router.Get("/hierarchies", hierarchies.Handler)

	router.Get("/download", download.Handler) // list high level dataset
	router.Get("/dataresources/{resourceId}", dataresource.Handler)

	// Geographic Hierarchies and Areas
	// http://localhost:20099/geoareas/?geoareatype=COUNTRY				// returns list of all geographical areas within an entity (e.g. Metropolitan Districts)
	router.Methods("GET").PathPrefix("/geoareas").Queries("geoareatype", "{type}").HandlerFunc(geoareatype.Handler)

	// http://localhost:20099/geoareas/?geohierarchy=2013ADMIN			// provide detailed information for a given geographic hierarchy
	// equivalent to http://localhost:20099/geohierarchies/2013ADMIN
	router.Methods("GET").PathPrefix("/geoareas").Queries("geohierarchy", "{hierarchy}").HandlerFunc(geohierarchy.Handler)

	// http://localhost:20099/geoareas/1000001
	router.Get("/geoareas/{geoareaId}", geoarea.Handler) // provide detailed information for a given geographic area

	// http://localhost:20099/geoareas
	router.Get("/geoareas", geoareas.Handler) // list all geographic areas across all hierarchies

	// http://localhost:20099/geohierarchies/2013ADMIN
	router.Get("/geohierarchies/{geohierarchyId}", geohierarchy.Handler) // provide detailed information for a given geographic hierarchy

	// http://localhost:20099/geohierarchies
	router.Get("/geohierarchies", geohierarchies.Handler) // list all geographic hierarchies

	log.Debug("Starting server", log.Data{"bind_addr": bindAddr})

	if err := http.ListenAndServe(bindAddr, alice); err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
}

func corsHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, req)
	})
}
