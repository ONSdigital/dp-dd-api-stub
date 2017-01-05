package geoareatype

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-dd-api-stub/models"
	"github.com/ONSdigital/go-ns/log"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	log.Namespace = "dp-dd-api-stub"

	log.Debug("Got query string:", log.Data{"geoareatype": req.URL.Query().Get("geoareatype")})

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

	geoAreas2011 := make([]*models.GeoArea, 2)

	geoAreas2011[0] = &models.GeoArea{
		ID:        "1000002",
		ExtCode:   "E92000001",
		Name:      "England",
		AreaType:  "COUNTRY",
		AreaLevel: "COUNTRY",
		Populations: populations,
	}

	geoAreas2011[1] = &models.GeoArea{
		ID:        "1000003",
		ExtCode:   "W92000004",
		Name:      "Wales",
		AreaType:  "COUNTRY",
		AreaLevel: "COUNTRY",
		Populations: populations,
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(geoAreas2011)

	w.WriteHeader(200)
}