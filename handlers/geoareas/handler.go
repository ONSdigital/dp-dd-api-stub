package geoareas

// Get All Entities from Hierarchy
// returns list of all geographical areas within an entity (e.g. Metropolitan Districts)

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-dd-api-stub/models"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	time_type := &models.TimeType{
		ID: "1",
	}

	time_period := &models.TimePeriod{
		ID:         "1",
		StartDate:  "2014-01-01",
		EndDate:    "2014-12-31",
		Name:       "2014",
		TimeTypeID: time_type.ID,
	}

	populations := make([]*models.Population, 1)

	populations[0] = &models.Population{
		ExtCode:      "K02000001",
		GeoAreaID:    "2000001",
		TimePeriodID: time_period.ID,
	}

	geoAreas := make([]*models.GeoArea, 6)

	geoAreas[0] = &models.GeoArea{
		ID:          "1000001",
		ExtCode:     "K04000001",
		Name:        "England and Wales",
		AreaType:    "EW",
		AreaLevel:   "EW",
		Populations: populations,
	}

	geoAreas[1] = &models.GeoArea{
		ID:          "1000002",
		ExtCode:     "E92000001",
		Name:        "England",
		AreaType:    "COUNTRY",
		AreaLevel:   "COUNTRY",
		Populations: populations,
	}

	geoAreas[2] = &models.GeoArea{
		ID:          "1000003",
		ExtCode:     "W92000004",
		Name:        "Wales",
		AreaType:    "COUNTRY",
		AreaLevel:   "COUNTRY",
		Populations: populations,
	}

	geoAreas[3] = &models.GeoArea{
		ID:          "2000001",
		ExtCode:     "K02000001",
		Name:        "United Kingdom",
		AreaType:    "UK",
		AreaLevel:   "UK",
		Populations: populations,
	}

	geoAreas[4] = &models.GeoArea{
		ID:          "2000002",
		ExtCode:     "K03000001",
		Name:        "Great Britain",
		RelGeoArea:  "2000001",
		AreaType:    "GB",
		AreaLevel:   "GB",
		Populations: populations,
	}

	geoAreas[5] = &models.GeoArea{
		ID:          "2000003",
		ExtCode:     "K04000001",
		Name:        "England and Wales",
		RelGeoArea:  "2000001",
		AreaType:    "EW",
		AreaLevel:   "EW",
		Populations: populations,
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(geoAreas)

	w.WriteHeader(200)
}
