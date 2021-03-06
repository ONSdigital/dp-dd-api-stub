package dimension

import (
	"errors"
	"github.com/ONSdigital/dp-dd-api-stub/stub"
	"github.com/ONSdigital/dp-dd-api-stub/models"
	"fmt"
	"os"
	"encoding/json"
)

func findDimension(datasetID string, edition string, version string,dimensionID string)  (dimension *models.Dimension, err error) {
	raw, err := stub.Asset("data/datasets/" + datasetID + "/" + edition + "/" + version + "/dimensions.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var dimensions []models.Dimension = []models.Dimension{}
	json.Unmarshal(raw, &dimensions)

	for _, dimension := range dimensions {
		if dimension.ID == dimensionID {
			return &dimension, nil
		}
	}
	return &models.Dimension{}, errors.New("Dimension id " + dimensionID + " not found.")
}

func legacyFindDimension(datasetID string, dimensionID string)  (dimension *models.Dimension, err error) {
	raw, err := stub.Asset("data/datasets/legacy/" + datasetID + "/dimensions.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var dimensions []models.Dimension = []models.Dimension{}
	json.Unmarshal(raw, &dimensions)

	for _, dimension := range dimensions {
		if dimension.ID == dimensionID {
			return &dimension, nil
		}
	}
	return &models.Dimension{}, errors.New("Dimension id " + dimensionID + " not found.")
}

func getDimensionHierarchy(dimensionID string) (hierarchy *models.Hierarchy, err error) {

	var hierarchyID string = "";
	switch dimensionID {
	case "SP00001": hierarchyID = "CI_000641"; break;
	case "T000111": hierarchyID = "TIME_001"; break;
	case "Geographic_Hierarchy": hierarchyID = "Geographic_Hierarchy"; break;
	}

	if (hierarchyID == "") {
		return nil, errors.New("Dimension id " + dimensionID + " not found.")
	}

	raw, err := stub.Asset("data/hierarchies/" + hierarchyID + ".json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	json.Unmarshal(raw, &hierarchy)
	hierarchy.ID = dimensionID;
	updateDimensionHierarchOptions(hierarchy.Options)
	return hierarchy, nil
}

func updateDimensionHierarchOptions(entries []*models.HierarchyEntry) {
	for _, option := range entries {
		option.HasData = true;
		if (len(option.Options) > 0) {
			updateDimensionHierarchOptions(option.Options)
		}
	}
}