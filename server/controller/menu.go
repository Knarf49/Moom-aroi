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

func GetMenuById(db *gorm.DB, id uint) ([]model.Dish, error) {
	var dish []model.Dish
	err := db.First(&dish, id).Error
	return dish, err
}

func GetRecommendMenu(db *gorm.DB) ([]model.Dish, error) {
	var dishes []model.Dish
	err := db.Where("Is_recommend = ?", true).Find(&dishes).Error
	return dishes, err
}
