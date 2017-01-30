package dimension

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-dd-api-stub/models"
	"github.com/ONSdigital/dp-dd-api-stub/stub"
	"fmt"
	"os"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	datasetID := req.URL.Query().Get(":datasetId")
	dimensionID := req.URL.Query().Get(":dimensionId")


	raw, err := stub.Asset("data/datasets/" + datasetID + "/dimensions.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var dimensions []models.Dimension = []models.Dimension{}
	json.Unmarshal(raw, &dimensions)

	dimension, err:= findDatasetIndexByID(dimensions, dimensionID)
	if err != nil {
		fmt.Println(err)
		_, err := json.Marshal(models.Dimension{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(400)
		//w.Write(response)
	} else {
		jsonEncoder := json.NewEncoder(w)
		jsonEncoder.Encode(dimension)
		w.WriteHeader(200)
	}

}
