package address

import (
	"context"

	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/address/models"
	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/address/repositories"
	"github.com/RuzimurodovDilshodbek/eater-service/src/infrastructure/utils"
	"gorm.io/gorm"
)

const (
	tableAddress = "eater.adress"
)

type addressRepoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repositories.AddressRepository {
	return &addressRepoImpl{
		db: db,
	}
}

func (r *addressRepoImpl) GetAddressTable(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Table(tableAddress)
}

func (r *addressRepoImpl) SaveAddress(ctx context.Context, address *models.Address) error {
	
	result := r.GetAddressTable(ctx).Create(&address)
	
	return result.Error

}

func (r *addressRepoImpl) UpdateAddress(ctx context.Context, address *models.Address) error {
	result := r.GetAddressTable(ctx).Save(&address)

	return result.Error
}

func (r *addressRepoImpl) DeleteAddress(ctx context.Context, addressId string) error {

	result := r.GetAddressTable(ctx).Delete("id = ?", addressId)

	return result.Error
}

func (r *addressRepoImpl) GetAddressById(ctx context.Context, addressId string) (*models.Address, error) {
	var address *models.Address
	result := r.GetAddressTable(ctx).First(address, addressId)

	if result.Error != nil {
		return nil, result.Error
	}

	return address, nil
}

func (r *addressRepoImpl) ListAddressByEater(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Address, error) {
	var addresses []*models.Address

	result := r.db.WithContext(ctx).Table(tableAddress).Where("eater_id = ?", eaterID)
	
	result.Scopes(utils.Paginate(page, pageSize), utils.Sort(sort)).Find(&addresses)

	if result.Error != nil {
		return nil, result.Error
	}
	return addresses, nil
}