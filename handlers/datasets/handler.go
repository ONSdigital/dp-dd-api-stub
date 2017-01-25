package datasets

import (
	"encoding/json"
	"github.com/ONSdigital/dp-dd-api-stub/models"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	var stubData models.Datasets = models.Datasets{}

	stubData.Items = []*models.Dataset{{
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
	}, {
		ID:               "CPI15",
		CustomerFacingID: "CPI15",
		Title:            "CPI15 Consumer Prices Index (COICOP).",
		URL:              "http://localhost:20099/datasets/CPI15",
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
	}, {
		ID:               "NO-META",
		CustomerFacingID: "NO-META",
		Title:            "Fake dataset with missing metadata field",
		URL:              "http://localhost:20099/datasets/NO-META",
	}, {
		ID:               "EMPTY-META",
		CustomerFacingID: "EMPTY-META",
		Title:            "Fake dataset with empty metadata",
		URL:              "http://localhost:20099/datasets/EMPTY-META",
		Metadata:         &models.Metadata{},
	}}

	stubData.Count = len(stubData.Items)
	stubData.StartIndex = 0
	stubData.ItemsPerPage = 10
	stubData.Total = len(stubData.Items)
	stubData.TotalPages = (len(stubData.Items) + 9) / stubData.ItemsPerPage
	stubData.Page = 1

	baseUrl := "http://" + req.Host + req.URL.EscapedPath()
	stubData.First = baseUrl + "?page=1"
	stubData.Last = baseUrl + "?page=" + strconv.Itoa(stubData.TotalPages)

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(stubData)

	w.WriteHeader(200)
}
