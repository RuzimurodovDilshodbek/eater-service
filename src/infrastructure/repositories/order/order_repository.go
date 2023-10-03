package order

import (
	"context"

	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/order/models"
	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/order/repositories"
	"gorm.io/gorm"
)

const (
	tableOrder = "order.order"
	tableCart  = "order.cart"
)

type orderRepoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repositories.OrderRepository {
	return &orderRepoImpl{
		db: db,
	}
}

func (r *orderRepoImpl) WithTx(ctx context.Context, f func(r repositories.OrderRepository) error) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		r := orderRepoImpl{
			db: tx,
		}
		if err := f(&r); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *orderRepoImpl) CreateOrder(ctx context.Context, order *models.Order) error {
	result := r.db.WithContext(ctx).Table(tableOrder).Create(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderRepoImpl) UpdateOrder(ctx context.Context, order *models.Order) error {
	result := r.db.WithContext(ctx).Table(tableOrder).Save(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderRepoImpl) GetOrderById(ctx context.Context, order_id string) (*models.Order, error) {
	var order models.Order
	result := r.db.WithContext(ctx).Table(tableOrder).First(&order, "id = ?", order_id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (r *orderRepoImpl) ListOrderByEaterId(ctx context.Context, eater_id string) ([]*models.Order, error) {
	var order []*models.Order
	result := r.db.WithContext(ctx).Table(tableOrder).Where(&order, "eater_id = ?", eater_id)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (r *orderRepoImpl) DeletoOrder(ctx context.Context, id string) error {
	var order models.Order
	result := r.db.WithContext(ctx).Table(tableOrder).Delete(&order, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
