package dataset

import (
	"encoding/json"
	"github.com/ONSdigital/dp-dd-api-stub/models"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	var stubData map[string]*models.Dataset = make(map[string]*models.Dataset)

	stubData["AF001EW"] = &models.Dataset{
		ID:               "AF001EW",
		CustomerFacingID: "AF001EW",
		Title:            "AF001EW  Members of the Armed Forces by residence type by sex by age",
		URL:              "http://localhost:20099/datasets/AF001EW",
		Metadata: &models.Metadata{
			Description:        "This dataset provides 2011 Census estimates that classify usual residents aged 16 and over who are members of the armed forces by residence type (household or communal resident), by sex and by age. The estimates are as at census day, 27 March 2011.",
			NationalStatistics: true,
			Contact: &models.Contact{
				Name:  "Alexa Bradley",
				Phone: "+44 (0)1329 444972",
				Email: "pop.info@ons.gsi.gov.uk",
			},
			ReleaseDate: "23 May 2014",
			NextRelease: "To be announced",
			Publications: []string{
				"http://localhost:20099/datasets/AF002",
				"http://localhost:20099/datasets/AF003",
				"http://localhost:20099/datasets/AF004",
			},
			TermsAndConditions: "N/A",
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
			Description:        "Consumer Price Index statistics by Time and Special Aggregate (type of economic activity). This dataset shows the movement of prices over the last five years within the UK economy, broken down by month and various classifications of economic activity. The economic classifications are derived from the COICOP (Classification Of Individual Consumption by Purpose) list.",
			NationalStatistics: true,
			Contact: &models.Contact{
				Name:  "James Tucker",
				Email: "cpi@ons.gsi.gov.uk",
				Phone: "+44 (0)1633 456900",
			},
			ReleaseDate: "17 January 2017",
			NextRelease: "14 February 2017",
			Methodology: []*models.Methodology{
				{Title: "Consumer Price Inflation (includes all 4 indices—CPI, CPIH, RPI and RPIJ)", URL: "https://www.ons.gov.uk/economy/inflationandpriceindices/qmis/consumerpriceinflationqmi"},
			},
			TermsAndConditions: "",
		},
		Dimensions: []*models.Dimension{
			{ID: "SP00001", Name: "Special aggregate"},
			{ID: "T000111", Name: "Time"},
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
