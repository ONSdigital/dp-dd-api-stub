package dimension

import (
	"encoding/json"
	"net/http"
	"github.com/ONSdigital/dp-dd-api-stub/models"
	"fmt"
)

func listHandler(w http.ResponseWriter, req *http.Request) {
	datasetID := req.URL.Query().Get(":datasetId")
	dimensionID := req.URL.Query().Get(":dimensionId")
	edition := req.URL.Query().Get(":edition")
	version := req.URL.Query().Get(":version")

	dimension, err := findDimension(datasetID, edition, version, dimensionID)
	if err != nil {
		fmt.Println(err)
		response, _ := json.Marshal(models.Dimension{})
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	w.WriteHeader(200)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(dimension)
}
