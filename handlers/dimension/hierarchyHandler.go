package dimension

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/ONSdigital/dp-dd-api-stub/models"
)

func hierarchyHandler(w http.ResponseWriter, req *http.Request) {
	//datasetID := req.URL.Query().Get(":datasetId")
	dimensionID := req.URL.Query().Get(":dimensionId")

	hierarchy, err := getDimensionHierarchy(dimensionID)
	if err != nil {
		fmt.Println(err)
		response, _ := json.Marshal(models.Dimension{})
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(hierarchy)
	w.WriteHeader(200)
}
