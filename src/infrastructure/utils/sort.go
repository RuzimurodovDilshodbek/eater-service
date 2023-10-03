package utils

import (
	"gorm.io/gorm"
	"github.com/RuzimurodovDilshodbek/eater-service/src/domain/order/models"

)

func Sort(sort string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		orderBy := "created_at DESC"
		if sort == models.SortByCreatedAtAsc {
			orderBy = "created_at ASC"
		}
		return db.Order(orderBy)
	}
}
