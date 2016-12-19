package geohierarchies

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-dd-api-stub/models"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	geoHierarchies := make([]*models.GeoHierarchy, 2)

	geoHierarchies[0] = &models.GeoHierarchy{
		Name: "2011GPH",
	}

	geoHierarchies[1] = &models.GeoHierarchy{
		Name: "2013ADMIN",
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(geoHierarchies)

	w.WriteHeader(200)
}
