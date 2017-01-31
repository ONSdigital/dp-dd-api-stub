package dimensions

import (
	"net/http"
	"github.com/ONSdigital/dp-dd-api-stub/stub"
	"fmt"
	"os"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	datasetID := req.URL.Query().Get(":datasetId")
	raw, err := stub.Asset("data/datasets/" + datasetID + "/dimensions.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	w.Write(raw)
	w.WriteHeader(200)
}
