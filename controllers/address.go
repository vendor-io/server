package controllers

import (
	"keyboardify-server/models"
	"keyboardify-server/models/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAddressesForUser(c echo.Context) error {
	uid := c.Param("uid")
	foundUser := FindUserViaUID(uid)

	var userAddresses []models.Address
	Db.Where("user_id = ?", foundUser.ID).Find(&userAddresses)

	for i, j := 0, len(userAddresses)-1; i < j; i, j = i+1, j-1 {
		userAddresses[i], userAddresses[j] = userAddresses[j], userAddresses[i]
	}

	return c.JSON(http.StatusOK, userAddresses)
}

func AddNewAddress(c echo.Context) error {
	var addAddress = new(dto.AddressDTO)

	if Err = c.Bind(addAddress); Err != nil {
		return Err
	}

	foundUser := FindUserViaUID(addAddress.UserID)

	var newAddress = models.Address{
		UserID:      foundUser.ID,
		FirstName:   addAddress.FirstName,
		LastName:    addAddress.LastName,
		Street:      addAddress.Street,
		HouseNumber: addAddress.HouseNumber,
		PostalCode:  addAddress.PostalCode,
		City:        addAddress.City,
		Country:     addAddress.Country,
		PhoneNumber: addAddress.PhoneNumber,
	}

	Db.Create(&newAddress)

	return c.JSON(http.StatusCreated, addAddress)
}
