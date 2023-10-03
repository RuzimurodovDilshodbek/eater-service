package services

import (
	"context"
	"time"

	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/order/models"
	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/order/repositories"
	"github.com/RuzimurodovDilshodbek/eater-service/src/infrastructure/rand"
)

type OrderService interface {
	SaveOrder(ctx context.Context, eaterID string, cart *models.Cart) (*models.Order, error)
	UpdateOrder(ctx context.Context, order *models.Order) (*models.Order, error) 
	DeleteOrder(ctx context.Context, orderID string) error
	GetOrderById(ctx context.Context, orderID string) (*models.Order, error)
	ListOrderByEaterId(ctx context.Context, eaterID string) ([]*models.Order, error)
}

type orderSvcImpl struct {
	orderRepo repositories.OrderRepository
}

func NewOrderService(
	orderRepo repositories.OrderRepository,
) OrderService {
	return &orderSvcImpl{
		orderRepo: orderRepo,
	}
}

func (s *orderSvcImpl) SaveOrder(ctx context.Context, eaterID string, cart *models.Cart) (*models.Order, error) {

	var (
	  now = time.Now()
	)
  
	restaurant := &models.RestaurantInfo{
	  Name:     cart.Restaurant.Name,
	  ImageUrl: cart.Restaurant.ImageUrl,
	}
  
	address := &models.AddressInfo{
	  ID:        cart.Delivery.AddressID,
	  Name:      cart.Delivery.Address.Name,
	  Longitude: cart.Delivery.Address.Longitude,
	  Latitude:  cart.Delivery.Address.Latitude,
	}
  
	delivery := &models.DeliveryInfo{
	  AddressID: address.ID,
	  Address:   address,
	  Time:      cart.Delivery.Time,
	  Notes:     cart.Delivery.Notes,
	}
  
	payment := &models.PaymentInfo{
	  Method:        cart.Payment.Method,
	  CardID:        cart.Payment.CardID,
	  DeliveryPrice: cart.Payment.DeliveryPrice,
	  TotalPrice:    cart.Payment.TotalPrice,
	}
	orderItems := []*models.OrderItem{}
  
	for key, b := range cart.Items {
  
	  orderItems[key] = &models.OrderItem{
		ID:         rand.UUID(),
		ProductID:  b.ProductID,
		Price:      b.ProductPrice,
		Quantity:   b.Quantity,
		TotalPrice: b.ProductPrice + b.Quantity,
		CreatedAt:  now,
	  }
	}
  
	order := &models.Order{
	  ID:            rand.UUID(),
	  EaterID:       eaterID,
	  Instruction:   cart.Instruction,
	  RestaurantID:  cart.RestaurantID,
	  Restaurant:    restaurant,
	  Delivery:      delivery,
	  Items:         orderItems,
	  Payment:       payment,
	  PaymentStatus: models.PaymentStatusPending,
	  Status:        models.OrderStatusPending,
	  CreatedAt:     now,
	  UpdatedAt:     now,
	}
  
	err := s.orderRepo.WithTx(ctx, func(r repositories.OrderRepository) error {
  
	
  
	  if err := r.CreateOrder(ctx, order); err != nil {
		return err
	  }
	  return nil
	})
  
	if err != nil {
	  return nil, err
	}
  
	return order, nil
  }


func (s *orderSvcImpl) UpdateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
		err := s.orderRepo.UpdateOrder(ctx, order)
		if err != nil {
			return nil, err
		}
		return order, nil
}

func (s *orderSvcImpl) DeleteOrder(ctx context.Context, orderID string) error {
		if err := s.orderRepo.DeletoOrder(ctx, orderID); err != nil {
			return err
		}
		return nil
}

func (s *orderSvcImpl) GetOrderById(ctx context.Context, orderID string) (*models.Order, error) {
	order, err := s.orderRepo.GetOrderById(ctx, orderID)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderSvcImpl) ListOrderByEaterId(ctx context.Context, eaterID string) ([]*models.Order, error) {
	addresses, err := s.orderRepo.ListOrderByEaterId(ctx, eaterID)

	if err != nil {
		return nil, err
	}

	return addresses, nil
}