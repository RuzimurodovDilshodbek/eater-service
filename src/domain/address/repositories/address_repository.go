package repositories

import (
	"context"

	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/address/models"
)

type AddressRepository interface {
	SaveAddress(ctx context.Context, address *models.Address) error
	UpdateAddress(ctx context.Context, address *models.Address) error
	DeleteAddress(ctx context.Context, addressId string) error
	GetAddressById(ctx context.Context, addressId string) (*models.Address, error)
	ListAddressByEater(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Address, error)
}
