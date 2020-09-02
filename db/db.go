package db

import (
	models "code/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		return nil
	}

	db.AutoMigrate(&models.User{})

	return db
}
