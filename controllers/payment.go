package controllers

import (
	"fmt"
	"keyboardify-server/models/dto"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

func PostCharge(c echo.Context) error {
	var chargeDTO = new(dto.PaymentDTO)

	if Err = c.Bind(chargeDTO); Err != nil {
		return Err
	}

	args := &stripe.ChargeParams{
		Amount:      stripe.Int64(int64(chargeDTO.Amount)),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String(fmt.Sprintf("Transaction for user %s", chargeDTO.ReceiptEmail)),
	}
	args.SetSource(chargeDTO.Source)

	ch, err := charge.New(args)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Request couldn't be processed - please try again.")
	}

	return c.JSON(http.StatusAccepted, ch)
}

func PostPaymentIntent(c echo.Context) error {
	var paymentDTO = new(dto.PaymentDTO)

	if Err = c.Bind(paymentDTO); Err != nil {
		return Err
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(paymentDTO.Amount)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	result, err := paymentintent.New(params)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Request couldn't be processed - please try again.")
	}

	data := dto.CheckoutData{
		ClientSecret: result.ClientSecret,
	}

	return c.JSON(http.StatusAccepted, data)
}
