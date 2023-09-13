package services

import (
	"context"
	"time"

	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/rating/models"
	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/rating/repositories"
	"go.uber.org/zap"
)

type RatingService interface {
	RateRestaurant(ctx context.Context, ID, EaterID, RestaurantID, Comment string, Rating int) (*models.RestaurantRating, error)
	UpdateRatingRestaurant(ctx context.Context, ID, EaterID, RestaurantID, Comment string, Rating int) (*models.RestaurantRating, error)
	ListRatingByEaterId(ctx context.Context, eaterID string) ([]*models.RestaurantRating, error)

	RateDelivery(ctx context.Context, ID, EaterID, OrderID, Comment string, Rating int) (*models.DeliveryRating, error)
	UpdateRatingDelivery(ctx context.Context, ID, EaterID, OrderID, Comment string, Rating int) (*models.DeliveryRating, error)
	GetDeliveryByOrderId(ctx context.Context, orderID string) (*models.DeliveryRating, error)
}

type ratingSvcImpl struct {
	ratingRepo repositories.RatingRepository
	logger     *zap.Logger
}

func NewRatingService(
	ratingRepo repositories.RatingRepository,
	logger *zap.Logger,
) RatingService {
	return &ratingSvcImpl{
		ratingRepo: ratingRepo,
		logger:     logger,
	}
}

func (s *ratingSvcImpl) RateRestaurant(ctx context.Context, ID, EaterID, RestaurantID, Comment string, Rating int) (*models.RestaurantRating, error) {
	rating := &models.RestaurantRating{
		ID:           ID,
		EaterID:      EaterID,
		RestaurantID: RestaurantID,
		Rating:       Rating,
		Comment:      Comment,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

	err := s.ratingRepo.WithTx(ctx, func(r repositories.RatingRepository) error {
		if err := r.RateRestaurant(ctx, rating); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return rating, nil
}

func (s *ratingSvcImpl) UpdateRatingRestaurant(ctx context.Context, ID, EaterID, RestaurantID, Comment string, Rating int) (*models.RestaurantRating, error) {
	rating := &models.RestaurantRating{
		ID:           ID,
		EaterID:      EaterID,
		RestaurantID: RestaurantID,
		Rating:       Rating,
		Comment:      Comment,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

	err := s.ratingRepo.WithTx(ctx, func(r repositories.RatingRepository) error {
		if err := r.UpdateRatingRestaurant(ctx, rating); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return rating, nil
}

func (s *ratingSvcImpl) ListRatingByEaterId(ctx context.Context, eaterID string) ([]*models.RestaurantRating, error) {
	rating, err := s.ListRatingByEaterId(ctx, eaterID)

	if err != nil {
		return nil, err
	}

	return rating, nil
}

func (s *ratingSvcImpl) RateDelivery(ctx context.Context, ID, EaterID, OrderID, Comment string, Rating int) (*models.DeliveryRating, error) {
	rating := &models.DeliveryRating{
		ID:        ID,
		EaterID:   EaterID,
		OrderID:   OrderID,
		Rating:    Rating,
		Comment:   Comment,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := s.ratingRepo.WithTx(ctx, func(r repositories.RatingRepository) error {
		if err := r.RateDelivery(ctx, rating); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return rating, nil
}

func (s *ratingSvcImpl) UpdateRatingDelivery(ctx context.Context, ID, EaterID, OrderID, Comment string, Rating int) (*models.DeliveryRating, error) {
	rating := &models.DeliveryRating{
		ID:        ID,
		EaterID:   EaterID,
		OrderID:   OrderID,
		Rating:    Rating,
		Comment:   Comment,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := s.ratingRepo.WithTx(ctx, func(r repositories.RatingRepository) error {
		if err := r.UpdateRatingDelivery(ctx, rating); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return rating, nil
}

func (s *ratingSvcImpl) GetDeliveryByOrderId(ctx context.Context, orderID string) (*models.DeliveryRating, error) {
	rating, err := s.ratingRepo.GetDeliveryByOrderId(ctx, orderID)

	if err != nil {
		return nil, err
	}

	return rating, nil
}
