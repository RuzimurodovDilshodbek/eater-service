package services

import (
	"context"
	"errors"

	dtos "github.com/RuzimurodovDilshodbek/eater-service/src/application/dtos/rating"
	ratingSvc "github.com/RuzimurodovDilshodbek/eater-service/src/domain/rating/services"
)

type RatingApplicationService interface {
	RateRestaurant(ctx context.Context, ID, EaterID, RestaurantID, Comment string, Rating int) (*dtos.SaveRestaurantRatingResponse, error)
	UpdateRatingRestaurant(ctx context.Context, ID, EaterID, RestaurantID, Comment string, Rating int) (*dtos.SaveRestaurantRatingResponse, error)
	ListRatingByEaterId(ctx context.Context, eaterID string) ([]*dtos.SaveRestaurantRatingResponse, error)

	RateDelivery(ctx context.Context, ID, EaterID, OrderID, Comment string, Rating int) (*dtos.SaveDeliveryRatingResponse, error)
	UpdateRatingDelivery(ctx context.Context, ID, EaterID, OrderID, Comment string, Rating int) (*dtos.SaveDeliveryRatingResponse, error)
	GetDeliveryByOrderId(ctx context.Context, orderID string) (*dtos.SaveDeliveryRatingResponse, error)
}

type ratingAppSvcImpl struct {
	ratingSvc ratingSvc.RatingService
}

func NewRatingApplicationService(
	ratingSvc ratingSvc.RatingService,
) RatingApplicationService {
	return &ratingAppSvcImpl{
		ratingSvc: ratingSvc,
	}
}

func (s *ratingAppSvcImpl) RateRestaurant(ctx context.Context, ID, EaterID, RestaurantID, Comment string, Rating int) (*dtos.SaveRestaurantRatingResponse, error) {
	if ID == "" {
		return nil, errors.New("invalid or missing Id")
	}
	if EaterID == "" {
		return nil, errors.New("invalid or missing EaterID")
	}
	if RestaurantID == "" {
		return nil, errors.New("invalid or missing RestaurantID")
	}
	if Comment == "" {
		return nil, errors.New("invalid or missing Comment")
	}
	if Rating == 0 {
		return nil, errors.New("invalid or missing Rating")
	}

	result, err := s.RateRestaurant(ctx, ID, EaterID, RestaurantID, Comment, Rating)
	if err != nil {
		return nil, err
	}

	return dtos.NewSaveRestaurantRatingResponse(result), nil
}

func (s *ratingAppSvcImpl) UpdateRatingRestaurant(ctx context.Context, ID, EaterID, RestaurantID, Comment string, Rating int) (*dtos.SaveRestaurantRatingResponse, error) {
	if ID == "" {
		return nil, errors.New("invalid or missing Id")
	}
	if EaterID == "" {
		return nil, errors.New("invalid or missing EaterID")
	}
	if RestaurantID == "" {
		return nil, errors.New("invalid or missing RestaurantID")
	}
	if Comment == "" {
		return nil, errors.New("invalid or missing Comment")
	}
	if Rating == 0 {
		return nil, errors.New("invalid or missing Rating")
	}

	result, err := s.UpdateRatingRestaurant(ctx, ID, EaterID, RestaurantID, Comment, Rating)
	if err != nil {
		return nil, err
	}

	return dtos.NewSaveRestaurantRatingResponse(&result), nil
}

func (s *ratingAppSvcImpl) ListRatingByEaterId(ctx context.Context, eaterID string) ([]*dtos.SaveRestaurantRatingResponse, error) {
	if eaterID == "" {
		return nil, errors.New("invalid or missing Rating")
	}

	result, err := s.UpdateRatingRestaurant(ctx, EaterID)
	if err != nil {
		return nil, err
	}

	return []dtos.NewSaveRestaurantRatingResponse(result), nil
}

func (s *ratingAppSvcImpl) RateDelivery(ctx context.Context, ID, EaterID, OrderID, Comment string, Rating int) (*dtos.SaveDeliveryRatingResponse, error) {
	if ID == "" {
		return nil, errors.New("invalid or missing ID")
	}
	if EaterID == "" {
		return nil, errors.New("invalid or missing EaterID")
	}
	if OrderID == "" {
		return nil, errors.New("invalid or missing OrderID")
	}
	if Comment == "" {
		return nil, errors.New("invalid or missing Comment")
	}
	if Rating == 0 {
		return nil, errors.New("invalid or missing Rating")
	}


	result, err := s.RateDelivery(ctx, ID, EaterID, OrderID, Comment, Rating)
	if err != nil {
		return nil, err
	}

	return dtos.SaveDeliveryRatingResponse(result), nil
}
 
func (s *ratingAppSvcImpl) UpdateRatingDelivery(ctx context.Context, ID, EaterID, OrderID, Comment string, Rating int) (*dtos.SaveDeliveryRatingResponse, error) {
	if ID == "" {
		return nil, errors.New("invalid or missing ID")
	}
	if EaterID == "" {
		return nil, errors.New("invalid or missing EaterID")
	}
	if OrderID == "" {
		return nil, errors.New("invalid or missing OrderID")
	}
	if Comment == "" {
		return nil, errors.New("invalid or missing Comment")
	}
	if Rating == 0 {
		return nil, errors.New("invalid or missing Rating")
	}


	result, err := s.UpdateRatingDelivery(ctx, ID, EaterID, OrderID, Comment, Rating)
	if err != nil {
		return nil, err
	}

	return dtos.SaveDeliveryRatingResponse(result), nil
}

func (s *ratingAppSvcImpl) GetDeliveryByOrderId(ctx context.Context, orderID string) (*dtos.SaveDeliveryRatingResponse, error) {
	if orderID == "" {
		return nil, errors.New("invalid or missing OrderID")
	}

	result, err := s.GetDeliveryByOrderId(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return dtos.SaveDeliveryRatingResponse(result), nil
}