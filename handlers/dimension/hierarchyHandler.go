package dimension

import (
	"encoding/json"
	"net/http"
	"fmt"
)

func hierarchyHandler(w http.ResponseWriter, req *http.Request) {
	//datasetID := req.URL.Query().Get(":datasetId")
	dimensionID := req.URL.Query().Get(":dimensionId")

	hierarchy, err := getDimensionHierarchy(dimensionID)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(404)
		return
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(hierarchy)
	w.WriteHeader(200)
}
