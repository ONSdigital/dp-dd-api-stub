package dimension

import (
	"github.com/ONSdigital/dp-dd-api-stub/models"
	"errors"
)

func findDatasetIndexByID(dimensions []models.Dimension, dimensionID string)  (dimension *models.Dimension, err error) {
	for _, dimension := range dimensions {
		if dimension.ID == dimensionID {
			return &dimension, nil
		}
	}
	return &models.Dimension{}, errors.New("Dimension id " + dimensionID + " not found.")
}