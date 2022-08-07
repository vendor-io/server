package controllers

import (
	"keyboardify-server/models"
	"keyboardify-server/models/dto"

	"gorm.io/gorm"
)

func FindUserViaUID(uid any) (user models.User) {
	Db.Where("uid = ?", uid).First(&user)

	return user
}

func CartDTOResolver(ids []uint, cartId uint) dto.CartWithTotalPriceDTO {
	var totalPrice uint = 0
	for i := range ids {
		var temporaryProductFromCart models.Product
		var productToExtractAmount models.CartProduct

		Db.Where("cart_id = ? AND product_id = ?", cartId, ids[i]).First(&productToExtractAmount)
		Db.Where("id = ?", ids[i]).First(&temporaryProductFromCart)

		totalPrice += temporaryProductFromCart.Price * productToExtractAmount.Amount
	}

	var productsToDTO []models.Product
	Db.Where(ids).Find(&productsToDTO)

	var productsDTO []dto.ProductInCartDTO
	for i := range productsToDTO {
		var categoryToDTO models.Category
		Db.Where("id = ?", productsToDTO[i].CategoryID).First(&categoryToDTO)

		var productToExtractAmount models.CartProduct
		Db.Where("cart_id = ? AND product_id = ?", cartId, productsToDTO[i].ID).First(&productToExtractAmount)

		var newProductToCartDTO = dto.ProductInCartDTO{
			ID:           productsToDTO[i].ID,
			Name:         productsToDTO[i].Name,
			MainImage:    productsToDTO[i].MainImage,
			Price:        productsToDTO[i].Price,
			Amount:       productToExtractAmount.Amount,
			TotalPrice:   productToExtractAmount.Amount * productsToDTO[i].Price,
			CategoryName: categoryToDTO.Name,
			CategorySlug: categoryToDTO.Slug,
		}

		productsDTO = append(productsDTO, newProductToCartDTO)
	}

	cart := dto.CartWithTotalPriceDTO{
		Products:   productsDTO,
		TotalPrice: totalPrice,
	}

	return cart
}

func ControllerDetailsResolver(cartDto dto.CartProductDTO) (user models.User, product models.Product, cart models.Cart, result *gorm.DB) {
	Db.Where("uid = ?", cartDto.UserID).First(&user)

	Db.Where("id = ?", cartDto.ProductID).First(&product)

	result = Db.Where("user_id = ?", user.ID).First(&cart)

	return user, product, cart, result
}
