package dataset

import (
	"encoding/json"
	"github.com/carlhembrough/dp-api-spike/models"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	var stubData *models.Dataset

	dimensions := make([]*models.Dimension, 4)

	dimensions[0] = &models.Dimension{
		ID:   "D000125",
		Name: "Sex",
	}
	dimensions[1] = &models.Dimension{
		ID:   "D000124",
		Name: "Residence Type",
	}
	dimensions[2] = &models.Dimension{
		ID:   "D000123",
		Name: "Age",
	}
	dimensions[3] = &models.Dimension{
		ID:   "D000126",
		Name: "2011 Statistical Geography Hierarchy",
	}

	stubData = &models.Dataset{
		ID:    "AF001EW",
		Title: "AF001EW  Members of the Armed Forces by residence type by sex by age",
		URL:   "http://localhost:20099/datasets/AF001EW",
		Metadata: &models.Metadata{
			Description: "This dataset provides 2011 Census estimates that classify usual residents aged 16 and over who are members of the armed forces by residence type (household or communal resident), by sex and by age. The estimates are as at census day, 27 March 2011.",
		},
		Dimensions: dimensions,
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(stubData)

	w.WriteHeader(200)
}
