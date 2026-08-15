package controller

import (
	model "server/schema"

	"gorm.io/gorm"
)

func GetAllMenu(db *gorm.DB) ([]model.Dish, error) {
	var dishes []model.Dish
	err := db.Find(&dishes).Error
	return dishes, err
}
