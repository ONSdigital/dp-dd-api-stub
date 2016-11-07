package dimension

import (
	"encoding/json"
	"github.com/carlhembrough/dp-api-spike/models"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	ageOptions := make([]*models.DimensionOption, 5)
	ageOptions[0] = &models.DimensionOption{
		ID:   "DO000265",
		Name: "All categories: Age 16 and over",
	}
	ageOptions[1] = &models.DimensionOption{
		ID:   "DO000266",
		Name: "Age 16 to 24",
	}
	ageOptions[2] = &models.DimensionOption{
		ID:   "DO000267",
		Name: "Age 25 to 34",
	}
	ageOptions[3] = &models.DimensionOption{
		ID:   "DO000268",
		Name: "Age 35 to 49",
	}
	ageOptions[4] = &models.DimensionOption{
		ID:   "DO000269",
		Name: "Age 50 and over",
	}

	dimension := &models.Dimension{
		ID:      "D000123",
		Name:    "Age",
		Options: ageOptions,
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(dimension)

	w.WriteHeader(200)
}
