package controller

import (
	models "server/schema"

	"gorm.io/gorm"
)

func SoftDeleteOrder(db *gorm.DB, orderId uint) error {
	tx := db.Select("Order_items").Delete(&models.Order{}, orderId)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
