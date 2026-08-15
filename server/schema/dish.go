package models

import (
	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	Title        string `gorm:"not null"`
	Desc         string
	Price        float64 `gorm:"not null"`
	Is_available bool    `gorm:"default:true"`
}