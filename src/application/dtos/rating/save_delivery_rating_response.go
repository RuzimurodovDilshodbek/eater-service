package dtos

import (
	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/rating/models"
)

type SaveDeliveryRatingResponse struct {
	DeliveryRating *models.DeliveryRating
}

func NewSaveDeliveryRatingResponse(deliveryRating *models.DeliveryRating) *SaveDeliveryRatingResponse {
	return &SaveDeliveryRatingResponse{
		DeliveryRating: deliveryRating,
	}
}
