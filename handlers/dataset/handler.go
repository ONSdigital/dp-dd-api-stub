package dataset

import (
	"encoding/json"
	"github.com/ONSdigital/dp-dd-api-stub/models"
	"net/http"
	"github.com/ONSdigital/dp-dd-api-stub/stub"
	"fmt"
	"os"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	datasetID := req.URL.Query().Get(":datasetId")
	raw, err := stub.Asset("data/datasets/" + datasetID + ".json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var stubData models.Dataset = models.Dataset{}
	json.Unmarshal(raw, &stubData)

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(stubData)

	w.WriteHeader(200)
}
