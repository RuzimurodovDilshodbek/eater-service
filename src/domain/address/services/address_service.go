package services

import (
	"context"
	"time"

	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/address/models"
	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/address/repositories"
	"github.com/RuzimurodovDilshodbek/eater-service/src/infrastructure/rand"
	"go.uber.org/zap"
)

type AddressService interface {
	CreateAddress(ctx context.Context, addressID, EaterID, name string, Latitude, Longitude float64) (*models.Address, error)
	UpdateAddress(ctx context.Context, addressID, EaterID, name string, Latitude, Longitude float64) (*models.Address, error)
	DeleteAddress(ctx context.Context, addressID string) error
	GetAddressById(ctx context.Context, addressID string) (*models.Address, error)
	ListAddressByEaterId(ctx context.Context, eaterID string) ([]*models.Address, error)
}

type addressSvcImpl struct {
	addressRepo repositories.AddressRepository
}

func NewAddressService(addressRepo repositories.AddressRepository, logger *zap.Logger) AddressService {
	return &addressSvcImpl{
		addressRepo: addressRepo,
	}
}

func (s *addressSvcImpl) CreateAddress(ctx context.Context, addressID, EaterID, name string, Latitude, Longitude float64) (*models.Address, error) {
	location := &models.Location{
		Longitude: Longitude,
		Latitude:  Latitude,
	}

	address := &models.Address{
		ID:        rand.UUID(),
		EaterID:   EaterID,
		Name:      name,
		Location:  location,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := s.addressRepo.SaveAddress(ctx, address)

	if err != nil {
		return nil, err
	}

	return address, nil
}

func (s *addressSvcImpl) UpdateAddress(ctx context.Context, addressID, EaterID, name string, Latitude, Longitude float64) (*models.Address, error) {
	location := &models.Location{
		Longitude: Longitude,
		Latitude:  Latitude,
	}

	address := &models.Address{
		ID:        addressID,
		EaterID:   EaterID,
		Name:      name,
		Location:  location,
		UpdatedAt: time.Now().UTC(),
	}

	err := s.addressRepo.UpdateAddress(ctx, address)

	if err != nil {
		return nil, err
	}

	return address, nil
}

func (s *addressSvcImpl) DeleteAddress(ctx context.Context, addressId string) error {

	return s.addressRepo.DeleteAddress(ctx, addressId)
}

func (s *addressSvcImpl) GetAddressById(ctx context.Context, addressID string) (*models.Address, error) {
	address, err := s.addressRepo.GetAddressById(ctx, addressID)

	if err != nil {
		return nil, err
	}

	return address, nil
}

func (s *addressSvcImpl) ListAddressByEaterId(ctx context.Context, eaterID string) ([]*models.Address, error) {
	addresses, err := s.addressRepo.ListAddressByEater(ctx, eaterID, sort, page, pageSize)
	if err != nil {
		return nil, err
	}

	return addresses, nil
}
