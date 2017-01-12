package datasets

import (
	"encoding/json"
	"github.com/ONSdigital/dp-dd-api-stub/models"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	var stubData models.Datasets = models.Datasets{}

	stubData.Items = []*models.Dataset{{
		ID:    "AF001EW",
		Title: "AF001EW  Members of the Armed Forces by residence type by sex by age",
		URL:   "http://localhost:20099/datasets/AF001EW",
		Metadata: &models.Metadata{
			Description: "This dataset provides 2011 Census estimates that classify usual residents aged 16 and over who are members of the armed forces by residence type (household or communal resident), by sex and by age. The estimates are as at census day, 27 March 2011.",
		},
	}, {
		ID:    "CPI15",
		Title: "CPI15 Consumer Prices Index (COICOP).",
		URL:   "http://localhost:20099/datasets/CPI15",
		Metadata: &models.Metadata{
			Description: "Consumer Price Index statistics by Time and Special Aggregate (type of economic activity). This dataset shows the movement of prices over the last five years within the UK economy, broken down by month and various classifications of economic activity. The economic classifications are derived from the COICOP (Classification Of Individual Consumption by Purpose) list.",
		},
	}}

	stubData.Count = len(stubData.Items)
	stubData.StartIndex = 0
	stubData.ItemsPerPage = 10
	stubData.Total = len(stubData.Items)

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(stubData)

	w.WriteHeader(200)
}
