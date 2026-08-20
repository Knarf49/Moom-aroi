package models

import (
	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	Url          string
	Title        string `gorm:"not null"`
	Desc         string
	Price        float64 `gorm:"not null"`
	Is_available bool    `gorm:"default:true"`
	Is_recommend bool    `gorm:"default:false"`
}
