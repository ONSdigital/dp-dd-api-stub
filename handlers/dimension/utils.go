package dimension

import (
	"github.com/ONSdigital/dp-dd-api-stub/models"
	"errors"
	"github.com/ONSdigital/dp-dd-api-stub/stub"
	"fmt"
	"os"
	"encoding/json"
)

func findDatasetIndexByID(dimensions []models.Dimension, dimensionID string)  (dimension *models.Dimension, err error) {
	for _, dimension := range dimensions {
		if dimension.ID == dimensionID {
			return &dimension, nil
		}
	}
	return &models.Dimension{}, errors.New("Dimension id " + dimensionID + " not found.")
}

func findHierarchyByDimensionID(dimensionID string) (hierarchy *models.Hierarchy, err error) {
	var hierarchyID string = "";
	switch dimensionID {
	case "SP00001": hierarchyID = ""; break;
	case "T000111": hierarchyID = ""; break;
	}

	if (hierarchyID == "") {
		return nil, errors.New("Dimension id " + dimensionID + " not found.")
	}

	raw, err := stub.Asset("data/hierarchy/" + hierarchyID + ".json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	json.Unmarshal(raw, &hierarchy)
	return hierarchy, nil
}
