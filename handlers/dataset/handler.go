package dataset

import (
	"encoding/json"
	"net/http"
	"github.com/ONSdigital/dp-dd-api-stub/models"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	var stubData map[string]*models.Dataset = make(map[string]*models.Dataset)

	stubData["AF001EW"] = &models.Dataset{
		ID:    "AF001EW",
		Title: "AF001EW  Members of the Armed Forces by residence type by sex by age",
		URL:   "http://localhost:20099/datasets/AF001EW",
		Metadata: &models.Metadata{
			Description: "This dataset provides 2011 Census estimates that classify usual residents aged 16 and over who are members of the armed forces by residence type (household or communal resident), by sex and by age. The estimates are as at census day, 27 March 2011.",
		},
		Dimensions: []*models.Dimension{
			{ID: "D000125", Name: "Sex"},
			{ID: "D000124", Name: "Residence Type"},
			{ID: "D000123", Name: "Age"},
			{ID: "D000126", Name: "2011 Statistical Geography Hierarchy"},
		},
	}

	stubData["CPI15"] = &models.Dataset{
		ID:    "CPI15",
		Title: "CPI15 Consumer Prices Index (COICOP).",
		URL:   "http://localhost:20099/datasets/CPI15",
		Metadata: &models.Metadata{
			Description: "This dataset provides 2011 Census estimates that classify usual residents aged 16 and over who are members of the armed forces by residence type (household or communal resident), by sex and by age. The estimates are as at census day, 27 March 2011.",
		},
		Dimensions: []*models.Dimension{
			{ID:   "SP00001", Name: "Special aggregate"},
			{ID:   "T000111", Name: "Time"},
		},
	}

	datasetID := req.URL.Query().Get(":datasetId")
	if _, ok := stubData[datasetID]; !ok {
		w.WriteHeader(404)
		return
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(stubData[datasetID])

	w.WriteHeader(200)
}
