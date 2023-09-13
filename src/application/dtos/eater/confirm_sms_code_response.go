package dtos

import "github.com/RuzimurodovDilshodbek/eater-service/src/domain/eater/models"

type ConfirmSMSCodeResponse struct {
	Token   string               `json:"token"`
	Profile *models.EaterProfile `json:"profile"`
}

func NewComfirmSMSCodeResponse(token string, profile *models.EaterProfile) *ConfirmSMSCodeResponse {
	return &ConfirmSMSCodeResponse{
		Profile: profile,
		Token:   token,
	}
}
 