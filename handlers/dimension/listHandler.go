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

	dimension, err := findDimension(datasetID, dimensionID)
	if err != nil {
		fmt.Println(err)
		response, _ := json.Marshal(models.Dimension{})
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(dimension)
	w.WriteHeader(200)
}
