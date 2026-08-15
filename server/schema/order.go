package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Status  string `gorm:"default:pending"` //'pending', 'preparing', 'completed'
	TableID uint   `gorm:"column:table_id;not null"`
	Table   Table  `gorm:"foreignKey:TableID"`
}

type Order_items struct {
	gorm.Model
	Quantity int   `gorm:"not null;default:1"`
	OrderID  uint  `gorm:"column:order_id;not null"`
	Order    Order `gorm:"foreignKey:OrderID"`
	DishID   uint  `gorm:"column:dish_id;not null"`
	Dish     Dish  `gorm:"foreignKey:DishID"`
}
