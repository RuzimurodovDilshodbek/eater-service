package repositories

import (
	"context"

	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/rating/models"
)

type RatingRepository interface {
	WithTx(ctx context.Context, f func(r RatingRepository) error) error
	RateRestaurant(ctx context.Context, rating *models.RestaurantRating) error
	UpdateRatingRestaurant(ctx context.Context, rating *models.RestaurantRating) error
	ListRatingByEaterId(ctx context.Context, eaterID string) ([]*models.RestaurantRating, error)

	RateDelivery(ctx context.Context, delivery *models.DeliveryRating) error
	UpdateRatingDelivery(ctx context.Context, delivery *models.DeliveryRating) error
	GetDeliveryByOrderId(ctx context.Context, orderID string) (*models.DeliveryRating, error)
}
