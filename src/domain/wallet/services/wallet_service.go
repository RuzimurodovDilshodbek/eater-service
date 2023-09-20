package services

import (
	"context"
	"time"

	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/wallet/models"
	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/wallet/repositories"
	"github.com/RuzimurodovDilshodbek/eater-service/src/infrastructure/rand"
)

type WalletService interface{
	AddCard(ctx context.Context, eaterID,cardNumber,cardToken string) error
	GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error)
	DeleteCard(ctx context.Context, cardID string) error
	GetCardsByEater(ctx context.Context, eaterID string) ([]*models.PaymentCard, error)
}

type walletSvcImpl struct{
	cardRepo repositories.CardRepository
}

func NewWalletService(cardRepo repositories.CardRepository) WalletService{
	return &walletSvcImpl{
		cardRepo: cardRepo,
	}
}

func (s *walletSvcImpl) AddCard(ctx context.Context, eaterID,cardNumber,cardToken string) error{
	parmentCard := &models.PaymentCard{
		ID: rand.UUID(),
		EaterID: eaterID,
		Number: cardNumber,
		CardToken: cardToken,
		IsVerifed: false,
		CreatedAt: time.Now().UTC(),
	}
	return s.cardRepo.AddCard(ctx,parmentCard)
}

func (s *walletSvcImpl) GetCard(ctx context.Context,cardID string)(*models.PaymentCard,error){
	return s.cardRepo.GetCard(ctx,cardID)
}

func (s *walletSvcImpl) DeleteCard(ctx context.Context,cardID string)error{
	return s.cardRepo.DeleteCard(ctx,cardID)
}
func (s *walletSvcImpl) GetCardsByEater(ctx context.Context,eaterID string)([]*models.PaymentCard,error){
	return s.cardRepo.GetCardsByEater(ctx,eaterID)
}