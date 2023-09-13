package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/RuzimurodovDilshodbek/eater-service/src/application/dtos"
	eatersvc "github.com/RuzimurodovDilshodbek/eater-service/src/domain/eater/services"
	"github.com/RuzimurodovDilshodbek/eater-service/src/infrastructure/jwt"
	"github.com/RuzimurodovDilshodbek/eater-service/src/infrastructure/validator"
)

type EaterApplicationService interface {
	SignupEater(ctx context.Context, phoneNumber string) (*dtos.EaterSignupResponse, error)
	ConfirmSMSCode(ctx context.Context, eaterID, smsCode string) (*dtos.ConfirmSMSCodeResponse, error)
	UpdateEaterProfile(ctx context.Context, eaterID, name, imageUrl string) (*dtos.UpdateEaterProfileResponse, error)
	GetEaterProfile(ctx context.Context, eaterID string) (*dtos.GetEaterProfileResponse, error)
}

type eaterAppSvcImpl struct {
	eaterSvc eatersvc.EaterService
	tokenSvc jwt.Service
}

func NewEaterApplicationService(
	eaterSvc eatersvc.EaterService,
	tokenSvc jwt.Service,
) EaterApplicationService {
	return &eaterAppSvcImpl{
		eaterSvc: eaterSvc,
		tokenSvc: tokenSvc,
	}
}

func (s *eaterAppSvcImpl) SignupEater(ctx context.Context, phoneNumber string) (*dtos.EaterSignupResponse, error) {
	if !validator.ValidateUzPhoneNumber(phoneNumber) {
		return nil, errors.New("Invalid number. Must be off 998{code}{number} format")
	}

	eaterID, err := s.eaterSvc.SignupEater(ctx, phoneNumber)
	if err != nil {
		return nil, err
	}

	response := dtos.NewEaterSignupResponse(eaterID)
	return response, nil
}

func (s *eaterAppSvcImpl) ConfirmSMSCode(ctx context.Context, eaterID, smsCode string) (*dtos.ConfirmSMSCodeResponse, error) {
	if eaterID == "" {
		return nil, errors.New("invalid or missing eaterId")
	}

	if smsCode == "" {
		return nil, errors.New("invalid or missing code")
	}

	profile, err := s.eaterSvc.ConfirmSMSCode(ctx, eaterID, smsCode)
	if err != nil {
		return nil, err
	}

	token, err := s.tokenSvc.CreateToken(ctx, profile.eaterID)
	if err != nil {
		return nil, err
	}

	response := dtos.ConfirmSMSCodeResponse{
		Profile: profile,
		Token:   token,
	}
	return &response, nil

}

func (s *eaterAppSvcImpl) GetEaterProfile(ctx context.Context, eaterID string) (*dtos.GetEaterProfileResponse, error) {
	if eaterID == "" {
		return nil, errors.New("invalid or missing eaterId")
	}

	profile, err := s.eaterSvc.GetEaterProfile(ctx, eaterID)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetEaterProfileResponse(profile), nil
}

func (s *eaterAppSvcImpl) UpdateEaterProfile(ctx context.Context, eaterID string, name string, imageUrl string) (*dtos.UpdateEaterProfileResponse, error) {
	if name == "" {
		return nil, fmt.Errorf("Invalid or mising profile name: %s", name)
	}

	profile, err := s.eaterSvc.UpdateEaterProfile(ctx, eaterID, name, imageUrl)
	if err != nil {
		return nil, err
	}

	return dtos.NewUpdateEaterProfileResponse(profile), nil

}
