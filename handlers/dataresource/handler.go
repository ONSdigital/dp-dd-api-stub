package dataresource

import (
	"encoding/json"
	"github.com/ONSdigital/dp-dd-api-stub/models"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	var stubData map[string]*models.DataResource = make(map[string]*models.DataResource)

	stubData["AF001EW"] = &models.DataResource{
		ID:    "AF001EW",
		Title: "AF001EW  Members of the Armed Forces by residence type by sex by age",
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
			Publications: []*models.Publication{
				{DatasetID: "AF002", Title: "AF002 Household Reference Persons (HRPs) who are members of the Armed Forces and associated persons in households with members of the Armed Forces by sex by age, local authorities in England and Wales"},
				{DatasetID: "AF003", Title: "AF003 Economic activity of associated people in households with members of the Armed Forces by sex, local authorities in England and Wales"},
				{DatasetID: "AF004", Title: "AF004 Members of the Armed Forces by workplace address by sex by age, local authorities in England and Wales"},
			},
		},
		Datasets: []string{"http://localhost:20099/datasets/AF001EW"},
	}

	stubData["CPI15"] = &models.DataResource{
		ID:    "CPI15",
		Title: "CPI15 Consumer Prices Index (COICOP).",
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
		},
		Datasets: []string{"http://localhost:20099/datasets/CPI15"},
	}

	datasetID := req.URL.Query().Get(":resourceId")
	if _, ok := stubData[datasetID]; !ok {
		w.WriteHeader(404)
		return
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(stubData[datasetID])

	w.WriteHeader(200)
}
