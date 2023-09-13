package dtos

import "github.com/RuzimurodovDilshodbek/eater-service/src/domain/address/models"

type ListAddressByEaterIdResponse struct {
	Address *models.Address 
}

func NewListAddressByEaterIdResponse(address *models.Address) *ListAddressByEaterIdResponse {
	return &ListAddressByEaterIdResponse{
		Address: address,
	}
}