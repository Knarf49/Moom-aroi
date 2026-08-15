package gorm

import (
	models "server/schema"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func SetupDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("aroi.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(models.AllModels...); err != nil {
		return nil, err
	}
	return db, nil
}
