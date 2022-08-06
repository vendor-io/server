package controllers

import (
	"fmt"
	"net/http"

	"keyboardify-server/models"
	"keyboardify-server/models/dto"

	"github.com/labstack/echo/v4"
)

func GetCartForUser(c echo.Context) error {
	uid := c.Param("uid")

	var foundUser models.User
	Db.Where("uid = ?", uid).First(&foundUser)

	var foundCart models.Cart
	result := Db.Where("user_id = ?", foundUser.ID).First(&foundCart)

	if result.Error != nil {
		return c.String(http.StatusOK, "Your cart is empty.")
	}

	var cartProducts []models.CartProduct
	Db.Where("cart_id = ?", foundCart.ID).Find(&cartProducts)

	if len(cartProducts) == 0 {
		cart := dto.CartWithTotalPriceDTO{
			Products:   []models.Product{},
			TotalPrice: 0,
		}

		return c.JSON(http.StatusOK, cart)
	}

	var cartProductsIDs []uint
	for i := 0; i <= len(cartProducts)-1; i++ {
		cartProductsIDs = append(cartProductsIDs, cartProducts[i].ProductID)
	}

	var totalPrice uint64 = 0
	for i := 0; i <= len(cartProductsIDs)-1; i++ {
		var temporaryProductFromCart models.Product
		Db.Where("id = ?", cartProductsIDs[i]).First(&temporaryProductFromCart)

		totalPrice += temporaryProductFromCart.Price
	}

	var productsToDTO []models.Product
	Db.Where(cartProductsIDs).Find(&productsToDTO)

	cart := dto.CartWithTotalPriceDTO{
		Products:   productsToDTO,
		TotalPrice: totalPrice,
	}

	return c.JSON(http.StatusOK, cart)
}

func AddProductToCart(c echo.Context) error {
	var addProduct = new(dto.CartProductDTO)

	if Err = c.Bind(addProduct); Err != nil {
		return Err
	}

	var foundUser models.User
	Db.Where("uid = ?", addProduct.UserID).First(&foundUser)

	var foundProduct models.Product
	Db.Where("id = ?", addProduct.ProductID).First(&foundProduct)

	var foundCart models.Cart
	result := Db.Where("user_id = ?", foundUser.ID).First(&foundCart)

	if result.Error != nil {
		fmt.Println("Creating a new cart.")

		newCartProduct := models.CartProduct{
			ProductID: foundProduct.ID,
			CartID:    foundCart.ID,
		}
		Db.Create(newCartProduct)

		var productsToDTO []models.Product
		Db.Where(newCartProduct.ProductID).Find(&productsToDTO)

		return c.JSON(http.StatusOK, productsToDTO)
	}

	if result.Error == nil {
		fmt.Println("Updating a cart.")

		newCartProduct := models.CartProduct{
			ProductID: foundProduct.ID,
			CartID:    foundCart.ID,
		}
		Db.Create(newCartProduct)

		var cartProducts []models.CartProduct
		Db.Where("cart_id = ?", foundCart.ID).Find(&cartProducts)

		var cartProductsIDs []uint
		for i := 0; i <= len(cartProducts)-1; i++ {
			cartProductsIDs = append(cartProductsIDs, cartProducts[i].ProductID)
		}

		var productsToDTO []models.Product
		Db.Where(cartProductsIDs).Find(&productsToDTO)

		cart := dto.CartDTO{
			Products: productsToDTO,
		}

		return c.JSON(http.StatusOK, cart)
	}

	return c.String(http.StatusBadRequest, "Request is invalid.")
}

func RemoveProductFromCart(c echo.Context) error {
	var productToRemove = new(dto.CartProductDTO)

	if Err = c.Bind(productToRemove); Err != nil {
		return Err
	}

	var foundUser models.User
	Db.Where("uid = ?", productToRemove.UserID).First(&foundUser)

	var foundProduct models.Product
	Db.Where("id = ?", productToRemove.ProductID).First(&foundProduct)

	var foundCart models.Cart
	Db.Where("user_id = ?", foundUser.ID).First(&foundCart)

	Db.Where("product_id = ? AND cart_id = ?", foundProduct.ID, foundCart.ID).Delete(&models.CartProduct{})

	return c.String(http.StatusAccepted, "Product has been removed from your cart.")
}
