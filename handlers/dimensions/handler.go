package dimensions

import (
	"net/http"
	"github.com/ONSdigital/dp-dd-api-stub/stub"
	"fmt"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	datasetID := req.URL.Query().Get(":datasetId")
	edition := req.URL.Query().Get(":edition")
	version := req.URL.Query().Get(":version")
	raw, err := stub.Asset("data/datasets/" + datasetID + "/" + edition + "/" + version + "/dimensions.json")
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(404)
		return
	}

	w.Write(raw)
	w.WriteHeader(200)
}
