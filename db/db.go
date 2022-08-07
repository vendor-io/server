package db

import (
	"fmt"
	"os"
	"path/filepath"

	"keyboardify-server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var ImagesPublicPath = filepath.Join("public", "images")

func Init() {
	db, err := gorm.Open(sqlite.Open("keyboardify_gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.ProductInOrder{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Cart{})
	db.AutoMigrate(&models.CartProduct{})
	db.AutoMigrate(&models.Address{})
	db.AutoMigrate(&models.AddressInOrder{})
	db.AutoMigrate(&models.Order{})

	os.MkdirAll(ImagesPublicPath, os.ModePerm)

}
