package services

import (
	"context"
	"time"

	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/order/models"
	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/order/repositories"
)

type OrderService interface {
	CreateAddress(
		ctx context.Context,
		Instruction,
		EaterID,
		RestaurantID,
		Status,
		PaymentStatus,
		OrderItemID,
		ProductID,
		Name string,
		Quantity,
		Price,
		TotalPrice int,
	) (*models.Order, error)
	UpdateAddress(
		ctx context.Context,
		Instruction,
		EaterID,
		RestaurantID,
		Status,
		PaymentStatus,
		OrderItemID,
		ProductID,
		Name string,
		Quantity,
		Price,
		TotalPrice int,
	) (*models.Order, error)
	DeleteAddress(ctx context.Context, orderID string) error
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

func (s *orderSvcImpl) CreateOrder(
	ctx context.Context,
	Instruction,
	EaterID,
	RestaurantID,
	Status,
	PaymentStatus,
	OrderItemID,
	ProductID,
	Name string,
	Quantity,
	Price,
	TotalPrice int,
) (*models.Order, error) {
	orderItem[] := models.OrderItem{
		ID:         OrderItemID,
		ProductID:  ProductID,
		Name:       Name,
		Quantity:   Quantity,
		Price:      Price,
		TotalPrice: TotalPrice,
		CreatedAt:  time.Now().UTC(),
	}

	order := &models.Order{
		EaterID:       EaterID,
		Instruction:   Instruction,
		RestaurantID:  RestaurantID,
		Items:         orderItem,
		Status:        Status,
		PaymentStatus: PaymentStatus,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}

	err := s.orderRepo.WithTx(ctx, func(r *repositories.OrderRepository) error {
		if err := r.CreateOrder(ctx, order); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (s *orderSvcImpl) UpdateOrder(
	ctx context.Context,
	Instruction,
	EaterID,
	RestaurantID,
	Status,
	PaymentStatus,
	OrderItemID,
	ProductID,
	Name string,
	Quantity,
	Price,
	TotalPrice int,
) (*models.Order, error) {
	orderItem := models.OrderItem{
		ID:         OrderItemID,
		ProductID:  ProductID,
		Name:       Name,
		Quantity:   Quantity,
		Price:      Price,
		TotalPrice: TotalPrice,
		CreatedAt:  time.Now().UTC(),
	}

	order := &models.Order{
		EaterID:       EaterID,
		Instruction:   Instruction,
		RestaurantID:  RestaurantID,
		Items:         orderItem,
		Status:        Status,
		PaymentStatus: PaymentStatus,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}

	err := s.orderRepo.WithTx(ctx, func(r *repositories.OrderRepository) error {
		if err := r.UpdateAddress(ctx, order); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (s *orderSvcImpl) DeleteOrder(ctx context.Context, orderID string) error {
	err := s.addressRepo.WithTx(ctx, func(r repositories.OrderRepository) error {
		if err := r.DeleteOrder(ctx, orderID); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
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

func (s *addressSvcImpl) ListOrderByEaterId(ctx context.Context, eaterID string) ([]*models.Order, error) {
	addresses, err := s.addressRepo.ListOrderByEaterId(ctx, eaterID)

	if err != nil {
		return nil, err
	}

	return addresses, nil
}