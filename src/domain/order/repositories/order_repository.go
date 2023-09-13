package repositories

import (
	"context"

	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/order/models"
)

type OrderRepository interface {
	WithTx(ctx context.Context, f func(r OrderRepository) error) error
	CreateOrder(ctx context.Context, order *models.Order) error
	UpdateOrder(ctx context.Context, order *models.Order) error
	GetOrderById(ctx context.Context, order_id string) (*models.Order, error)
	ListOrderByEaterId(ctx context.Context, eater_id string) ([]*models.Order, error)
	DeleteAddress(ctx context.Context, id string) error
}
