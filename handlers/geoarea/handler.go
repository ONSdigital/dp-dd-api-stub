package geoarea

// Get Entities from Hierarchy - request by Entity,
// returns list of all geographical areas within an entity (e.g. Metropolitan Districts)


import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-dd-api-stub/models"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	time_type := &models.TimeType{
		ID:    "1",
	}

	time_period := &models.TimePeriod{
		ID:        "1",
		StartDate: "2014-01-01",
		EndDate:   "2014-12-31",
		Name:      "2014",
		TimeTypeID: time_type.ID,
	}

	populations := make([]*models.Population, 1)

	populations[0] = &models.Population{
		ExtCode:      "K02000001",
		GeoAreaID:    "2000001",
		TimePeriodID: time_period.ID,
	}


	geoAreas2011 := make([]*models.GeoArea, 1)

	geoAreas2011[0] = &models.GeoArea{
		ID:          "1000001",
		ExtCode:     "K04000001",
		Name:        "England and Wales",
		AreaType:    "EW",
		AreaLevel:   "EW",
		Populations: populations,
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(geoAreas2011)

	w.WriteHeader(200)
}