package controllers

import (
	"keyboardify-server/models"
	"keyboardify-server/models/dto"
)

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
