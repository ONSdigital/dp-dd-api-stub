package dimensions

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

	sexOptions := make([]*models.DimensionOption, 2)
	sexOptions[0] = &models.DimensionOption{
		ID:   "DO000153",
		Name: "Male",
	}
	sexOptions[1] = &models.DimensionOption{
		ID:   "DO000154",
		Name: "Female",
	}

	residenceTypeOptions := make([]*models.DimensionOption, 3)
	residenceTypeOptions[0] = &models.DimensionOption{
		ID:   "DO000161",
		Name: "All categories: Residence Type",
	}
	residenceTypeOptions[1] = &models.DimensionOption{
		ID:   "DO000162",
		Name: "Lives in a household",
	}
	residenceTypeOptions[2] = &models.DimensionOption{
		ID:   "DO000163",
		Name: "Lives in a communal establishment",
	}

	geographyOptions := make([]*models.DimensionOption, 3)
	geographyOptions[0] = &models.DimensionOption{
		ID:   "K04000001",
		Name: "England and Wales",
	}
	geographyOptions[1] = &models.DimensionOption{
		ID:   "E92000001",
		Name: "England",
	}
	geographyOptions[2] = &models.DimensionOption{
		ID:   "W92000004",
		Name: "Wales",
	}

	dimensions := make([]*models.Dimension, 4)

	dimensions[0] = &models.Dimension{
		ID:      "D000125",
		Name:    "Sex",
		Options: sexOptions,
	}
	dimensions[1] = &models.Dimension{
		ID:      "D000124",
		Name:    "Residence Type",
		Options: residenceTypeOptions,
	}
	dimensions[2] = &models.Dimension{
		ID:      "D000123",
		Name:    "Age",
		Options: ageOptions,
	}
	dimensions[3] = &models.Dimension{
		ID:      "D000126",
		Name:    "2011 Statistical Geography Hierarchy",
		Options: geographyOptions,
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(dimensions)

	w.WriteHeader(200)
}
