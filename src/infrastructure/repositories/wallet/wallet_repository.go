package wallet

import (
	"context"

	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/wallet/models"
	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/wallet/repositories"
	"github.com/RuzimurodovDilshodbek/eater-service/src/infrastructure/repositories/wallet"
	"gorm.io/gorm"
)

const (
	cardTable = "eater.payment_cards"
)

type cardRepositoryImpl struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repositories.CardRepository{
	return &cardRepositoryImpl{
		db: db,
	}
}


func (r *cardRepositoryImpl) GetCardTable(ctx context.Context) *gorm.DB{
	return r.db.WithContext(ctx).Table(cardTable)
}

func (r *cardRepositoryImpl) AddCard(ctx context.Context,paymentCard *models.PaymentCard) error{
	return r.GetCardTable(ctx).Create(paymentCard).Error
}

func (r *cardRepositoryImpl) GetCard(ctx context.Context,cardID string) (*models.PaymentCard,error){
	
	var paymentCard *models.PaymentCard

	result := r.GetCardTable(ctx).First(&paymentCard,"id = ?",cardID)
	if result.Error != nil {
		return nil,result.Error		
	}

	return paymentCard,nil

}

func (r *cardRepositoryImpl) DeleteCard(ctx context.Context,cardID string) error{
	return r.GetCardTable(ctx).Delete("id = ?",cardID).Error
}

func (r *cardRepositoryImpl) GetCardsByEater(ctx context.Context,eaterID string) ([]*models.PaymentCard,error){
	
	var paymentCards []*models.PaymentCard
	result := r.GetCardTable(ctx).Where("eater_id =?",eaterID).Find(paymentCards)

	if result.Error != nil {
		return nil, result.Error
	}

	return paymentCards,nil
}