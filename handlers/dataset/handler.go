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
			&models.Dimension{ID: "D000125", Name: "Sex"},
			&models.Dimension{ID: "D000124", Name: "Residence Type"},
			&models.Dimension{ID: "D000123", Name: "Age"},
			&models.Dimension{ID: "D000126", Name: "2011 Statistical Geography Hierarchy"},
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
			&models.Dimension{
				ID:   "D000125",
				Name: "Special aggregate",
				Options: []*models.DimensionOption{
					&models.DimensionOption{
						ID:   "FOOD0001",
						Name: "(01) Food and non-alcoholic beverages",
						Options: []*models.DimensionOption{
							&models.DimensionOption{ID: "OPT02", Name: "Food"},
							&models.DimensionOption{ID: "OPT03", Name: "Bread and cereals"},
							&models.DimensionOption{ID: "OPT04", Name: "Meat"},
							&models.DimensionOption{ID: "OPT05", Name: "Fish"},
							&models.DimensionOption{ID: "OPT06", Name: "Milk, cheese and eggs"},
							&models.DimensionOption{ID: "OPT07", Name: "Oils and fats"},
							&models.DimensionOption{ID: "OPT08", Name: "Fruit"},
						},
					},
					&models.DimensionOption{
						ID:   "HEALTH0002",
						Name: "(02) Health",
						Options: []*models.DimensionOption{
							&models.DimensionOption{ID: "OPT02", Name: "Medical products, appliances and equipment"},
							&models.DimensionOption{ID: "OPT03", Name: "Pharmaceutical products"},
							&models.DimensionOption{ID: "OPT04", Name: "Other medical and therapeutic equipment"},
							&models.DimensionOption{ID: "OPT05", Name: "Out-patient services"},
							&models.DimensionOption{ID: "OPT06", Name: "Medical services and paramedical services"},
							&models.DimensionOption{ID: "OPT07", Name: "Dental services"},
							&models.DimensionOption{ID: "OPT08", Name: "In-patient service"},
							&models.DimensionOption{ID: "OPT09", Name: "Medical and paramedic services"},
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
