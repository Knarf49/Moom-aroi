package models

import (
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	Label string `gorm:"not null"`
}
