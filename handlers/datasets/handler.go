package datasets

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-dd-api-stub/models"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	var stubData models.Datasets = models.Datasets{}

	stubData.Items = make([]*models.Dataset, 3)

	stubData.Items[0] = &models.Dataset{
		ID:    "AF001EW",
		Title: "AF001EW  Members of the Armed Forces by residence type by sex by age",
		URL:   "http://localhost:20099/datasets/AF001EW",
		Metadata: &models.Metadata{
			Description: "This dataset provides 2011 Census estimates that classify usual residents aged 16 and over who are members of the armed forces by residence type (household or communal resident), by sex and by age. The estimates are as at census day, 27 March 2011.",
		},
	}

	stubData.Items[1] = &models.Dataset{
		ID:    "G39",
		Title: "House Price Statistics for Small Areas",
		URL:   "http://localhost:20099/datasets/G39",
		Metadata: &models.Metadata{
			Description: "These house price statistics are an accurate representation of the price paid for residential properties sold in a given area. They are useful for assessing the affordability of housing in small areas as well as broad patterns in prices and the number of house sales over time. Time Series Start Date: Jan-95 to Dec-95 End Date: Jan-13 to Dec-13 Periodicity Calendar Year Latest Coverage England, Wales Source Office for National Statistics Last Updated 17 February 2015",
		},
	}

	stubData.Items[2] = &models.Dataset{
		ID:    "CPI15",
		Title: "CPI15 Consumer Prices Index (COICOP).",
		URL:   "http://localhost:20099/datasets/CPI15",
		Metadata: &models.Metadata{
			Description: "Consumer Price Index statistics by Time and Special Aggregate (type of economic activity). This dataset shows the movement of prices over the last five years within the UK economy, broken down by month and various classifications of economic activity. The economic classifications are derived from the COICOP (Classification Of Individual Consumption by Purpose) list.",
		},
	}

	stubData.Count = len(stubData.Items)
	stubData.StartIndex = 0
	stubData.ItemsPerPage = 10
	stubData.Total = len(stubData.Items)

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(stubData)

	w.WriteHeader(200)
}
