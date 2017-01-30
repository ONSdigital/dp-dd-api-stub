package hierarchies

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-dd-api-stub/models"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	var hierarchiesStub []*models.Hierarchy = []*models.Hierarchy{{
		ID:   "Time",
		Name: "Time",
		Type: "time",
	}, {
		ID:   "CI_000641",
		Name: "Special Aggregate",
		Type: "classification",
	}}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(hierarchiesStub)

	w.WriteHeader(200)
}
