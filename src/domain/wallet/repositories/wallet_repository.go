package repositories

import (
	"context"

	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/wallet/models"
)

type CardRepository interface {
	AddCard(ctx context.Context, paymentCard *models.PaymentCard) error
	GetCard(ctx context.Context, cardD string) (*models.PaymentCard, error)
	DeleteCard(ctx context.Context, cardD string) error
	GetCardsByEater(ctx context.Context, eaterID string) ([]*models.PaymentCard, error)
}
