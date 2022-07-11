package db

import (
	"fmt"

	"keyboardify-server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {
	db, err := gorm.Open(sqlite.Open("keyboardify_gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Order{}, &models.Product{})
}
