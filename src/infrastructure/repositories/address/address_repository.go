package address

import (
	"context"

	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/address/models"
	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/address/repositories"
	"gorm.io/gorm"
)

const (
	tableAddress = "adress.adress"
)

type addressRepoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repositories.AddressRepository {
	return &addressRepoImpl{
		db: db,
	}
}

func (r *addressRepoImpl) WithTx(ctx context.Context, f func(r repositories.AddressRepository) error) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		r := addressRepoImpl{
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

func (r *addressRepoImpl) CreateAddress(ctx context.Context, address *models.Address) error {
	result := r.db.WithContext(ctx).Table(tableAddress).Create(address)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *addressRepoImpl) UpdateAddress(ctx context.Context, address *models.Address) error {
	result := r.db.WithContext(ctx).Table(tableAddress).Save(address)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *addressRepoImpl) DeleteAddress(ctx context.Context, id string) error {
	var address models.Address
	result := r.db.WithContext(ctx).Table(tableAddress).Delete(&address, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *addressRepoImpl) GetAddressById(ctx context.Context, id string) (*models.Address, error) {
	var address models.Address
	result := r.db.WithContext(ctx).Table(tableAddress).First(&address, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &address, nil
}

func (r *addressRepoImpl) ListAddressByEaterId(ctx context.Context, eaterID string) ([]*models.Address, error) {
	var address []*models.Address
	result := r.db.WithContext(ctx).Table(tableAddress).Where(address, "eater_id = ?", eaterID)
	if result.Error != nil {
		return nil, result.Error
	}
	return address, nil
}
