package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"keyboardify-server/db"
	"keyboardify-server/models"
	"keyboardify-server/models/dto"

	"github.com/labstack/echo/v4"
)

func GetAllProducts(c echo.Context) error {
	var products []models.Product
	Db.Find(&products)

	return c.JSON(http.StatusOK, products)
}

func GetProductById(c echo.Context) error {
	id := c.Param("id")
	product := new(models.Product)
	Db.Where("ID = ?", id).First(&product)
	return c.JSON(http.StatusOK, product)
}

func AddNewProduct(c echo.Context) error {
	var p = dto.ProductDTO{
		Name:        c.FormValue("productName"),
		EAN:         c.FormValue("productEan"),
		Price:       c.FormValue("productPrice"),
		Description: c.FormValue("productDescription"),
		Category:    c.FormValue("productCategory"),
	}

	// Product Images
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["productImages"]
	var images []string

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(filepath.Join(db.ImagesPublicPath, file.Filename))
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		images = append(images, fmt.Sprintf("%s/api/public/images/%s", os.Getenv("URL"), file.Filename))
	}

	imagesToString := strings.Join(images, ";")

	var foundCategory models.Category
	Db.Where("Name = ?", p.Category).First(&foundCategory)

	eanUint, eanUintErr := strconv.ParseUint(p.EAN, 10, 32)
	priceUint, priceUintErr := strconv.ParseUint(p.Price, 10, 32)

	product := models.Product{
		Name:        p.Name,
		EAN:         eanUint,
		Price:       priceUint,
		Description: p.Description,
		Category:    foundCategory,
		Images:      imagesToString,
		Stock:       0,
	}

	println(eanUintErr, priceUintErr)

	result := Db.Where("EAN = ?", eanUint).First(&models.Product{})

	if result.Error != nil {
		Db.Create(&product)

		Db.Model(&models.Category{}).Where("Name = ?", foundCategory.Name).Update("ItemsAmount", foundCategory.ItemsAmount+1)

		return c.JSON(http.StatusCreated, product)
	}

	return c.JSON(http.StatusConflict, "Product already exists!")
}
