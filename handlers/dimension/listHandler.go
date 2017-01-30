package dimension

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-dd-api-stub/models"
)

func ViewHandler(w http.ResponseWriter, req *http.Request) {

	optionsEnglandNorthEast := make([]*models.DimensionOption, 4)
	optionsEnglandNorthEast[0] = &models.DimensionOption{
		ID:   "2000017",
		Name: "Hartlepool",
	}
	optionsEnglandNorthEast[1] = &models.DimensionOption{
		ID:   "2000018",
		Name: "Middlesbrough",
	}
	optionsEnglandNorthEast[2] = &models.DimensionOption{
		ID:   "2000019",
		Name: "Redcar and Cleveland",
	}
	optionsEnglandNorthEast[3] = &models.DimensionOption{
		ID:   "2000020",
		Name: "Stockton-on-Tees",
	}

	optionsEngland := make([]*models.DimensionOption, 4)
	optionsEngland[0] = &models.DimensionOption{
		ID:      "E12000001",
		Name:    "North East",
		Options: optionsEnglandNorthEast,
	}
	optionsEngland[1] = &models.DimensionOption{
		ID:   "E12000002",
		Name: "North West",
	}
	optionsEngland[2] = &models.DimensionOption{
		ID:   "E12000003",
		Name: "Yorkshire and The Humber",
	}
	optionsEngland[3] = &models.DimensionOption{
		ID:   "E12000004",
		Name: "East Midlands",
	}

	options2 := make([]*models.DimensionOption, 2)
	options2[0] = &models.DimensionOption{
		ID:      "E92000001",
		Name:    "England",
		Options: optionsEngland,
	}
	options2[1] = &models.DimensionOption{
		ID:   "W92000004",
		Name: "Wales",
	}

	options := make([]*models.DimensionOption, 1)
	options[0] = &models.DimensionOption{
		ID:      "K04000001",
		Name:    "England and Wales",
		Options: options2,
	}

	dimension := &models.Dimension{
		ID:      "D000123",
		Name:    "Geography",
		Options: options,
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(dimension)

	w.WriteHeader(200)
}
