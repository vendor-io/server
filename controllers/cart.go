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
			Products:   []dto.ProductInCartDTO{},
			TotalPrice: 0,
		}

		return c.JSON(http.StatusOK, cart)
	}

	var cartProductsIDs []uint
	for i := 0; i <= len(cartProducts)-1; i++ {
		cartProductsIDs = append(cartProductsIDs, cartProducts[i].ProductID)
	}

	cart := CartDTOResolver(cartProductsIDs, foundCart.ID)

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

		var newCartForUser = models.Cart{
			UserID:         foundUser.ID,
			ProductsInCart: []models.CartProduct{},
		}
		Db.Create(newCartForUser)

		newCartProduct := models.CartProduct{
			ProductID: foundProduct.ID,
			CartID:    newCartForUser.ID,
			Amount:    addProduct.Amount,
		}
		Db.Create(newCartProduct)

		productsIDs := []uint{foundCart.ID}
		cart := CartDTOResolver(productsIDs, newCartForUser.ID)

		return c.JSON(http.StatusOK, cart)
	}

	if result.Error == nil {
		fmt.Println("Updating a cart.")

		var cartProductToUpdate models.CartProduct
		result := Db.Where("cart_id = ? AND product_id = ?", foundCart.ID, foundProduct.ID).First(&cartProductToUpdate)

		if result.Error != nil {
			fmt.Println("A new product has been added - creating new CartProduct instance.")

			newCartProduct := models.CartProduct{
				ProductID: foundProduct.ID,
				CartID:    foundCart.ID,
				Amount:    addProduct.Amount,
			}
			Db.Create(newCartProduct)
		}

		if result.Error == nil {
			fmt.Println("A product is already in cart - updating amount.")

			Db.Model(&cartProductToUpdate).Where("cart_id = ? AND product_id = ?", foundCart.ID, foundProduct.ID).Update("amount", cartProductToUpdate.Amount+addProduct.Amount)
		}

		var cartProducts []models.CartProduct
		Db.Where("cart_id = ?", foundCart.ID).Find(&cartProducts)

		var cartProductsIDs []uint
		for i := range cartProducts {
			cartProductsIDs = append(cartProductsIDs, cartProducts[i].ProductID)
		}

		cart := CartDTOResolver(cartProductsIDs, foundCart.ID)

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

	var cartProducts []models.CartProduct
	Db.Where("cart_id = ?", foundCart.ID).Find(&cartProducts)

	if len(cartProducts) == 0 {
		cart := dto.CartWithTotalPriceDTO{
			Products:   []dto.ProductInCartDTO{},
			TotalPrice: 0,
		}

		return c.JSON(http.StatusOK, cart)
	}

	var cartProductsIDs []uint
	for i := 0; i <= len(cartProducts)-1; i++ {
		cartProductsIDs = append(cartProductsIDs, cartProducts[i].ProductID)
	}

	var totalPrice uint = 0
	for i := 0; i <= len(cartProductsIDs)-1; i++ {
		var temporaryProductFromCart models.Product
		Db.Where("id = ?", cartProductsIDs[i]).First(&temporaryProductFromCart)

		totalPrice += temporaryProductFromCart.Price
	}

	var productsToDTO []models.Product
	Db.Where(cartProductsIDs).Find(&productsToDTO)

	var productsDTO []dto.ProductInCartDTO
	for i := range productsToDTO {
		var categoryToDTO models.Category
		Db.Where("id = ?", productsToDTO[i].CategoryID).First(&categoryToDTO)

		var newProductToCartDTO = dto.ProductInCartDTO{
			ID:           productsToDTO[i].ID,
			Name:         productsToDTO[i].Name,
			MainImage:    productsToDTO[i].MainImage,
			Price:        productsToDTO[i].Price,
			CategoryName: categoryToDTO.Name,
			CategorySlug: categoryToDTO.Slug,
		}

		productsDTO = append(productsDTO, newProductToCartDTO)
	}

	cart := dto.CartWithTotalPriceDTO{
		Products:   productsDTO,
		TotalPrice: totalPrice,
	}

	return c.JSON(http.StatusOK, cart)
}
