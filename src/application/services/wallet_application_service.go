package services

import (
	"context"
	"fmt"

	"github.com/RuzimurodovDilshodbek/eater-service/src/application/dtos"
	walletsvc "github.com/RuzimurodovDilshodbek/eater-service/src/domain/wallet/services"
)

type WalletApplicationService interface {
	AddCard(ctx context.Context, eaterId, number, cardToken, restaurant string) (*dtos.GetPaymentCardResponse, error)
	DeleteCard(ctx context.Context, cardId string) (*dtos.GetPaymentCardResponse, error)
	GetCard(ctx context.Context, cardId string) (*dtos.GetPaymentCardResponse, error)
	ListCardByEater(ctx context.Context, eaterID string, sort string, page, pageSize int) (*dtos.ListPaymentCardResponse, error)
}

type walletAppSvcImpl struct {
	walletSvc walletsvc.WalletService
}

func NewWalletApplicationService(walletSvc walletsvc.WalletService) *walletAppSvcImpl {
	return &walletAppSvcImpl{
		walletSvc: walletSvc,
	}
}

func (s *walletAppSvcImpl) AddCard(ctx context.Context, eaterId, number, cardToken, restaurant string) (*dtos.GetPaymentCardResponse, error) {

	if eaterId == "" {
		return nil, fmt.Errorf("invalid or missing eaterId: %s", eaterId)
	}
	if cardToken == "" {
		return nil, fmt.Errorf("Invalid or missing cardToken: %s", cardToken)
	}
	if number == "" {
		return nil, fmt.Errorf("Invalid or missing number: %s", number)
	}

	if restaurant == "" {
		return nil, fmt.Errorf("Invalid or missing restaurant: %s", restaurant)
	}
	paymentCard, err := s.walletSvc.AddCard(ctx, eaterId, number, cardToken, restaurant)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetPaymentCardResponse(paymentCard), nil
}

func (s *walletAppSvcImpl) DeleteCard(ctx context.Context, cardId string) (*dtos.GetPaymentCardResponse, error) {

	if cardId == "" {
		return nil, fmt.Errorf("Invalid or missing cardId: %s", cardId)
	}

	paymentCard, err := s.walletSvc.DeleteCard(ctx, cardId)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetPaymentCardResponse(paymentCard), nil
}

func (s *walletAppSvcImpl) GetCard(ctx context.Context, cardId string) (*dtos.GetPaymentCardResponse, error) {

	if cardId == "" {
		return nil, fmt.Errorf("Invalid or missing cardId: %s", cardId)
	}

	paymentCard, err := s.walletSvc.GetCard(ctx, cardId)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetPaymentCardResponse(paymentCard), nil
}

func (s *walletAppSvcImpl) ListCardByEater(ctx context.Context, eaterID string, sort string, page, pageSize int) (*dtos.ListPaymentCardResponse, error) {

	if eaterID == "" {
		return nil, fmt.Errorf("Invalid or missing eaterID: %s", eaterID)
	}

	paymentCards, err := s.walletSvc.ListCardByEater(ctx, eaterID, sort, page, pageSize)
	if err != nil {
		return nil, err
	}

	return dtos.NewListPaymentCardResponse(paymentCards), nil
}
