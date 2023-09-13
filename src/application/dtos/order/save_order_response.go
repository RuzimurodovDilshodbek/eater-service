package dtos

import "github.com/RuzimurodovDilshodbek/eater-service/src/domain/order/models"

type SaveOrderResponse struct {
	Order *models.Order
}

func NewSaveOrderResponse(order *models.Order) *SaveOrderResponse {
	return &SaveOrderResponse{
		Order: order,
	}
}
