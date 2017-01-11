package geohierarchy

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

	geoAreas2013 := make([]*models.GeoArea, 3)

	geoAreas2013[0] = &models.GeoArea{
		ID:          "2000001",
		ExtCode:     "K02000001",
		Name:        "United Kingdom",
		AreaType:    "UK",
		AreaLevel:   "UK",
		Populations: populations,
	}

	geoAreas2013[1] = &models.GeoArea{
		ID:          "2000002",
		ExtCode:     "K03000001",
		Name:        "Great Britain",
		RelGeoArea:  "2000001",
		AreaType:    "GB",
		AreaLevel:   "GB",
		Populations: populations,
	}

	geoAreas2013[2] = &models.GeoArea{
		ID:          "2000003",
		ExtCode:     "K04000001",
		Name:        "England and Wales",
		RelGeoArea:  "2000001",
		AreaType:    "EW",
		AreaLevel:   "EW",
		Populations: populations,
	}

	geoHierarchy := &models.GeoHierarchy{
		Name:     "2013ADMIN",
		GeoAreas: geoAreas2013,
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(geoHierarchy)

	w.WriteHeader(200)
}
