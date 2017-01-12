package dataset

import (
	"encoding/json"
	"net/http"
	"github.com/ONSdigital/dp-dd-api-stub/models"
	"fmt"
)

const S3URL_FMT = "s3://test-bucket/dir1/%s"

func Handler(w http.ResponseWriter, req *http.Request) {

	var stubData map[string]*models.Dataset = make(map[string]*models.Dataset)

	stubData["AF001EW"] = &models.Dataset{
		ID:    "AF001EW",
		S3URL: fmt.Sprintf(S3URL_FMT, "AF001EW"),
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
		S3URL: fmt.Sprintf(S3URL_FMT, "CPI15"),
		URL:   "http://localhost:20099/datasets/CPI15",
		Metadata: &models.Metadata{
			Description: "This dataset provides 2011 Census estimates that classify usual residents aged 16 and over who are members of the armed forces by residence type (household or communal resident), by sex and by age. The estimates are as at census day, 27 March 2011.",
		},
		Dimensions: []*models.Dimension{{
			ID:   "D000125",
			Name: "Special aggregate",
			Options: []*models.DimensionOption{{
				ID:   "FOOD0001",
				Name: "(01) Food and non-alcoholic beverages",
				Options: []*models.DimensionOption{
					{ID: "OPT02", Name: "Food"},
					{ID: "OPT03", Name: "Bread and cereals"},
					{ID: "OPT04", Name: "Meat"},
					{ID: "OPT05", Name: "Fish"},
					{ID: "OPT06", Name: "Milk, cheese and eggs"},
					{ID: "OPT07", Name: "Oils and fats"},
					{ID: "OPT08", Name: "Fruit"},
				},
			}, {
				ID:   "HEALTH0002",
				Name: "(02) Health",
				Options: []*models.DimensionOption{
					{ID: "OPT02", Name: "Medical products, appliances and equipment"},
					{ID: "OPT03", Name: "Pharmaceutical products"},
					{ID: "OPT04", Name: "Other medical and therapeutic equipment"},
					{ID: "OPT05", Name: "Out-patient services"},
					{ID: "OPT06", Name: "Medical services and paramedical services"},
					{ID: "OPT07", Name: "Dental services"},
					{ID: "OPT08", Name: "In-patient service"},
					{ID: "OPT09", Name: "Medical and paramedic services"},
				},
			},
			},
		},
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
