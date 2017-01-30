package hierarchies

import (
	"encoding/json"
	"fmt"
	"github.com/ONSdigital/dp-dd-api-stub/models"
	"github.com/ONSdigital/dp-dd-api-stub/stub"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	raw, err := stub.Asset("data/hierarchies.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var stubData []models.Hierarchy = []models.Hierarchy{}
	json.Unmarshal(raw, &stubData)

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(stubData)

	w.WriteHeader(200)
}
