package data

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-dd-api-stub/models"
	"fmt"
)

const S3URL_FMT = "s3://test-bucket/dir1/%s"

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

	if req.FormValue("D000125") == "DO000154" {
		rows := make([]*models.Row, 2)

		dimensions3 := make([]*models.DimensionOption, 4)
		dimensions3[0] = &models.DimensionOption{
			ID:   "DO000154",
			Name: "Female",
		}
		dimensions3[1] = &models.DimensionOption{
			ID:   "DO000162",
			Name: "Lives in a household",
		}
		dimensions3[2] = &models.DimensionOption{
			ID:   "DO000267",
			Name: "Age 25 to 34",
		}
		dimensions3[3] = &models.DimensionOption{
			ID:   "E92000001",
			Name: "England",
		}
		rows[0] = &models.Row{
			Observation: 6742,
			Dimensions:  dimensions3,
		}

		dimensions4 := make([]*models.DimensionOption, 4)
		dimensions4[0] = &models.DimensionOption{
			ID:   "DO000154",
			Name: "Female",
		}
		dimensions4[1] = &models.DimensionOption{
			ID:   "DO000162",
			Name: "Lives in a household",
		}
		dimensions4[2] = &models.DimensionOption{
			ID:   "DO000267",
			Name: "Age 25 to 34",
		}
		dimensions4[3] = &models.DimensionOption{
			ID:   "E92000001",
			Name: "England",
		}
		rows[1] = &models.Row{
			Observation: 432,
			Dimensions:  dimensions4,
		}

		table := &models.Table{
			Rows: rows,
		}

		stubData = &models.Dataset{
			ID:         "AF001EW",
			S3URL:      fmt.Sprintf(S3URL_FMT, "AF001EW"),
			Dimensions: dimensions,
			Data:       table,
		}
	} else {

		rows := make([]*models.Row, 5)

		dimensions0 := make([]*models.DimensionOption, 4)
		dimensions0[0] = &models.DimensionOption{
			ID:   "DO000153",
			Name: "Male",
		}
		dimensions0[1] = &models.DimensionOption{
			ID:   "DO000162",
			Name: "Lives in a household",
		}
		dimensions0[2] = &models.DimensionOption{
			ID:   "DO000267",
			Name: "Age 25 to 34",
		}
		dimensions0[3] = &models.DimensionOption{
			ID:   "K04000001",
			Name: "England and Wales",
		}
		rows[0] = &models.Row{
			Observation: 153223,
			Dimensions:  dimensions0,
		}

		dimensions1 := make([]*models.DimensionOption, 4)
		dimensions1[0] = &models.DimensionOption{
			ID:   "DO000153",
			Name: "Male",
		}
		dimensions1[1] = &models.DimensionOption{
			ID:   "DO000162",
			Name: "Lives in a household",
		}
		dimensions1[2] = &models.DimensionOption{
			ID:   "DO000267",
			Name: "Age 25 to 34",
		}
		dimensions1[3] = &models.DimensionOption{
			ID:   "E92000001",
			Name: "England",
		}
		rows[1] = &models.Row{
			Observation: 146348,
			Dimensions:  dimensions1,
		}

		dimensions2 := make([]*models.DimensionOption, 4)
		dimensions2[0] = &models.DimensionOption{
			ID:   "DO000153",
			Name: "Male",
		}
		dimensions2[1] = &models.DimensionOption{
			ID:   "DO000162",
			Name: "Lives in a household",
		}
		dimensions2[2] = &models.DimensionOption{
			ID:   "DO000267",
			Name: "Age 25 to 34",
		}
		dimensions2[3] = &models.DimensionOption{
			ID:   "W92000004",
			Name: "Wales",
		}
		rows[2] = &models.Row{
			Observation: 6875,
			Dimensions:  dimensions2,
		}

		dimensions3 := make([]*models.DimensionOption, 4)
		dimensions3[0] = &models.DimensionOption{
			ID:   "DO000154",
			Name: "Female",
		}
		dimensions3[1] = &models.DimensionOption{
			ID:   "DO000162",
			Name: "Lives in a household",
		}
		dimensions3[2] = &models.DimensionOption{
			ID:   "DO000267",
			Name: "Age 25 to 34",
		}
		dimensions3[3] = &models.DimensionOption{
			ID:   "E92000001",
			Name: "England",
		}
		rows[3] = &models.Row{
			Observation: 6742,
			Dimensions:  dimensions3,
		}

		dimensions4 := make([]*models.DimensionOption, 4)
		dimensions4[0] = &models.DimensionOption{
			ID:   "DO000154",
			Name: "Female",
		}
		dimensions4[1] = &models.DimensionOption{
			ID:   "DO000162",
			Name: "Lives in a household",
		}
		dimensions4[2] = &models.DimensionOption{
			ID:   "DO000267",
			Name: "Age 25 to 34",
		}
		dimensions4[3] = &models.DimensionOption{
			ID:   "W92000004",
			Name: "Wales",
		}
		rows[4] = &models.Row{
			Observation: 432,
			Dimensions:  dimensions4,
		}

		table := &models.Table{
			Rows: rows,
		}

		stubData = &models.Dataset{
			ID:         "AF001EW",
			S3URL:      fmt.Sprintf(S3URL_FMT, "AF001EW"),
			Dimensions: dimensions,
			Data:       table,
		}
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(stubData)

	w.WriteHeader(200)
}
