package services

import (
	"context"
	"errors"

	dtos "github.com/RuzimurodovDilshodbek/eater-service/src/application/dtos/order"
	orderSvc "github.com/RuzimurodovDilshodbek/eater-service/src/domain/order/services"
)

type OrderApplicationService interface {
	CreateOrder(
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
	) (*dtos.SaveOrderResponse, error)
	UpdateOrder(
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
	) (*dtos.SaveOrderResponse, error)
	DeleteOrder(ctx context.Context, orderID string) error
	GetOrderById(ctx context.Context, orderID string) (*dtos.SaveOrderResponse, error)
	ListOrderByEaterId(ctx context.Context, eaterID string) ([]*dtos.SaveOrderResponse, error)
}

type orderAppSvcImpl struct {
	orderSvc orderSvc.OrderService
}

func NewOrderApplicationService(
	orderSvc orderSvc.OrderService,
) OrderApplicationService {
	return &orderAppSvcImpl{
		orderSvc: orderSvc,
	}
}

func (s *orderAppSvcImpl) CreateOrder(
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
) (*dtos.SaveOrderResponse, error) {
	if Instruction == "" {
		return nil, errors.New("Instruction is is required!")
	}
	if EaterID == "" {
		return nil, errors.New("EaterID is is required!")
	}
	if RestaurantID == "" {
		return nil, errors.New("RestaurantID is is required!")
	}
	if PaymentStatus == "" {
		return nil, errors.New("PaymentStatus is is required!")
	}
	if OrderItemID == "" {
		return nil, errors.New("OrderItemID is is required!")
	}
	if ProductID == "" {
		return nil, errors.New("ProductID is is required!")
	}
	if Name == "" {
		return nil, errors.New("Name is is required!")
	}
	if Quantity == 0 {
		return nil, errors.New("Quantity is is required!")
	}
	if Price == 0 {
		return nil, errors.New("Price is is required!")
	}
	if TotalPrice == 0 {
		return nil, errors.New("TotalPrice is is required!")
	}

	result, err := s.CreateOrder(ctx, Instruction, EaterID, RestaurantID, Status, PaymentStatus, OrderItemID, ProductID, Name, Quantity, Price, TotalPrice)
	if err != nil {
		return nil, err
	}

	return dtos.SaveOrderResponse(result), nil
}

func (s *orderAppSvcImpl) UpdateOrder(
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
) (*dtos.SaveOrderResponse, error) {
	if Instruction == "" {
		return nil, errors.New("Instruction is is required!")
	}
	if EaterID == "" {
		return nil, errors.New("EaterID is is required!")
	}
	if RestaurantID == "" {
		return nil, errors.New("RestaurantID is is required!")
	}
	if PaymentStatus == "" {
		return nil, errors.New("PaymentStatus is is required!")
	}
	if OrderItemID == "" {
		return nil, errors.New("OrderItemID is is required!")
	}
	if ProductID == "" {
		return nil, errors.New("ProductID is is required!")
	}
	if Name == "" {
		return nil, errors.New("Name is is required!")
	}
	if Quantity == 0 {
		return nil, errors.New("Quantity is is required!")
	}
	if Price == 0 {
		return nil, errors.New("Price is is required!")
	}
	if TotalPrice == 0 {
		return nil, errors.New("TotalPrice is is required!")
	}

	result, err := s.UpdateOrder(ctx, Instruction, EaterID, RestaurantID, Status, PaymentStatus, OrderItemID, ProductID, Name, Quantity, Price, TotalPrice)
	if err != nil {
		return nil, err
	}

	return dtos.SaveOrderResponse(result), nil
}

func (s *orderAppSvcImpl) DeleteOrder(ctx context.Context, orderID string) error {
	if orderID == "" {
		return errors.New("orderID is is required!")
	}

	err := s.DeleteOrder(ctx, orderID)
	if err != nil {
		return err
	}

	return nil
}

func (s *orderAppSvcImpl) GetOrderById(ctx context.Context, orderID string) (*dtos.SaveOrderResponse, error) {
	if orderID == "" {
		return nil, errors.New("orderID is is required!")
	}

	result, err := s.GetOrderById(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return dtos.NewSaveOrderResponse(result), nil
}

func (s *orderAppSvcImpl) ListOrderByEaterId(ctx context.Context, eaterID string) ([]*dtos.SaveOrderResponse, error) {
	if eaterID == "" {
		return nil, errors.New("eaterID is is required!")
	}

	result, err := s.ListOrderByEaterId(ctx, eaterID)
	if err != nil {
		return nil, err
	}

	return dtos.NewSaveOrderResponse(result), nil
}
