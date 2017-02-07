package dataset

import (
	"net/http"
	"github.com/ONSdigital/dp-dd-api-stub/stub"
	"fmt"
	"strings"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	datasetID := req.URL.Query().Get(":datasetId")
	edition := req.URL.Query().Get(":edition")
	version := req.URL.Query().Get(":version")
	uriPath := req.URL.Path
	var r []byte
	println(uriPath)

	if edition == "" && version == "" {
		raw, err := stub.Asset("data/datasets/" + datasetID + "/editions.json" )
		println("I was here")
		if err != nil || strings.Contains(uriPath, "edition") || strings.Contains(uriPath, "value") {
			println("booooo")
			error404(w, err)
			println("bah")
			return
		}
		r = raw
	} else if version == "" {
		raw, err := stub.Asset("data/datasets/" + datasetID + "/" + edition + "/versions.json" )
		if err != nil || strings.Contains(uriPath, "value"){
			error404(w, err)
			return
		}
		r = raw
	} else {
		raw, err := stub.Asset("data/datasets/" + datasetID + "/" + edition + "/" + version + "/" + version + ".json" )
		if err != nil {
			error404(w, err)
			return
		}
		r = raw
	}

	w.Write(r)
	w.WriteHeader(200)
}

func LegacyHandler(w http.ResponseWriter, req *http.Request) {

	uuid := req.URL.Query().Get(":uuid")
	raw, err := stub.Asset("data/datasets/legacy/" + uuid + ".json")
	if err != nil {
		error404(w, err)
		return
	}

	w.Write(raw)
	w.WriteHeader(200)
}

func error404(w http.ResponseWriter, err error) {
	println("gonna be a error")
	fmt.Printf("%+v", err)
	println("just saying")
	w.WriteHeader(404)
}