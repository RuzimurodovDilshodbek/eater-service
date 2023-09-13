package rating

import (
	"context"

	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/rating/models"
	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/rating/repositories"
	"gorm.io/gorm"
)

const (
	tableRestaurantRating = "rating.restaurant_rating"
	tableDeliveryRating   = "rating.delivery_rating"
)

type ratingRepoImpl struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) repositories.RatingRepository {
	return &ratingRepoImpl{
		db: db,
	}
}

func (r *ratingRepoImpl) WithTx(ctx context.Context, f func(r repositories.RatingRepository) error) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		r := ratingRepoImpl{
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

func (r *ratingRepoImpl) RateRestaurant(ctx context.Context, rating *models.RestaurantRating) error {
	result := r.db.WithContext(ctx).Table(tableRestaurantRating).Create(rating)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ratingRepoImpl) UpdateRatingRestaurant(ctx context.Context, rating *models.RestaurantRating) error {
	result := r.db.WithContext(ctx).Table(tableRestaurantRating).Save(rating)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ratingRepoImpl) ListRatingByEaterId(ctx context.Context, eaterID string) ([]*models.RestaurantRating, error) {
	var rating []models.RestaurantRating
	result := r.db.WithContext(ctx).Table(tableRestaurantRating).Where(&rating, "eater_id = ?", eaterID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &rating, nil
}

func (r *ratingRepoImpl) RateDelivery(ctx context.Context, delivery *models.DeliveryRating) error {
	result := r.db.WithContext(ctx).Table(tableDeliveryRating).Create(delivery)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ratingRepoImpl) UpdateRatingDelivery(ctx context.Context, delivery *models.DeliveryRating) error {
	result := r.db.WithContext(ctx).Table(tableDeliveryRating).Save(delivery)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ratingRepoImpl) GetDeliveryByOrderId(ctx context.Context, orderID string) (*models.DeliveryRating, error) {
	var delivery models.DeliveryRating
	result := r.db.WithContext(ctx).Table(tableDeliveryRating).Where(&delivery, "order_id = ?", orderID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &delivery, nil
}
