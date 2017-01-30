package dimensions

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-dd-api-stub/models"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	var stubDimensions map[string][]*models.Dimension = make(map[string][]*models.Dimension)

	stubDimensions["AF001EW"] = []*models.Dimension{{
		ID:   "D000125",
		Name: "Sex",
		Options: []*models.DimensionOption{
			{ID: "DO000153", Name: "Male"},
			{ID: "DO000154", Name: "Female"},
		},
	}, {
		ID:   "D000124",
		Name: "Residence Type",
		Options: []*models.DimensionOption{
			{ID: "DO000161", Name: "All categories: Residence Type"},
			{ID: "DO000162", Name: "Lives in a household"},
			{ID: "DO000163", Name: "Lives in a communal establishment"},
		},
	}, {
		ID:   "D000123",
		Name: "Age",
		Options: []*models.DimensionOption{
			{ID: "DO000265", Name: "All categories: Age 16 and over"},
			{ID: "DO000266", Name: "Age 16 to 24"},
			{ID: "DO000267", Name: "Age 25 to 34"},
			{ID: "DO000268", Name: "Age 35 to 49"},
			{ID: "DO000269", Name: "Age 50 and over"},
		},
	}, {
		ID:   "D000126",
		Name: "2011 Statistical Geography Hierarchy",
		Type: "geography",
		Options: []*models.DimensionOption{
			{ID: "K04000001", Name: "England and Wales"},
			{ID: "E92000001", Name: "England"},
			{ID: "W92000004", Name: "Wales"},
		},
	}}

	stubDimensions["CPI15"] = []*models.Dimension{{
		ID:   "SP00001",
		Name: "Special aggregate",
		Type: "classification",
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
		}},
	}, {
		ID:   "T000111",
		Name: "Time",
		Type: "time",
		Options: []*models.DimensionOption{
			{ID: "OPT02", Name: "February 2013"},
			{ID: "OPT02", Name: "March 2013"},
			{ID: "OPT02", Name: "April 2013"},
			{ID: "OPT02", Name: "May 2013"},
			{ID: "OPT02", Name: "June 2013"},
			{ID: "OPT02", Name: "July 2013"},
			{ID: "OPT02", Name: "August 2013"},
			{ID: "OPT02", Name: "September 2013"},
			{ID: "OPT02", Name: "October 2013"},
			{ID: "OPT02", Name: "November 2013"},
			{ID: "OPT02", Name: "December 2013"},
			{ID: "OPT02", Name: "January 2014"},
			{ID: "OPT02", Name: "February 2014"},
			{ID: "OPT02", Name: "March 2014"},
			{ID: "OPT02", Name: "April 2014"},
			{ID: "OPT02", Name: "May 2014"},
			{ID: "OPT02", Name: "June 2014"},
			{ID: "OPT02", Name: "July 2014"},
			{ID: "OPT02", Name: "August 2014"},
			{ID: "OPT02", Name: "September 2014"},
			{ID: "OPT02", Name: "October 2014"},
			{ID: "OPT02", Name: "November 2014"},
			{ID: "OPT02", Name: "December 2014"},
			{ID: "OPT02", Name: "January 2015"},
			{ID: "OPT02", Name: "February 2015"},
			{ID: "OPT02", Name: "March 2015"},
			{ID: "OPT02", Name: "April 2015"},
			{ID: "OPT02", Name: "May 2015"},
			{ID: "OPT02", Name: "June 2015"},
			{ID: "OPT02", Name: "July 2015"},
			{ID: "OPT02", Name: "August 2015"},
			{ID: "OPT02", Name: "September 2015"},
			{ID: "OPT02", Name: "October 2015"},
			{ID: "OPT02", Name: "November 2015"},
			{ID: "OPT02", Name: "December 2015"},
			{ID: "OPT02", Name: "January 2016"},
			{ID: "OPT02", Name: "February 2016"},
			{ID: "OPT02", Name: "March 2016"},
			{ID: "OPT02", Name: "April 2016"},
			{ID: "OPT02", Name: "May 2016"},
			{ID: "OPT02", Name: "June 2016"},
			{ID: "OPT02", Name: "July 2016"},
			{ID: "OPT02", Name: "August 2016"},
			{ID: "OPT02", Name: "September 2016"},
			{ID: "OPT02", Name: "October 2016"},
		},
	}}

	datasetID := req.URL.Query().Get(":datasetId")
	if _, ok := stubDimensions[datasetID]; !ok {
		w.WriteHeader(404)
		return
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(stubDimensions[datasetID])

	w.WriteHeader(200)
}
