package models

import (
	"time"
)

// Order Information
type Order struct {
	ID              string    `json:"id"`
	UserID          string    `json:"userId"`
	Products        []Product `json:"products"`
	NumberOfDays    int       `json:"numberOfDays"`
	OrderValue      int       `json:"orderValue"`
	CashDeposit     int       `json:"cashDeposit"`
	PaymentMethod   string    `json:"paymentMethod"`
	DeliveryAddress string    `json:"deliveryAddress"`
	PlacedOn        time.Time `json:"placedOn"`
	DeliveredOn     time.Time `json:"deliveredOn"`
	PickedUpOn      time.Time `json:"pickedUpOn"`
	ReturnedOn      time.Time `json:"returnedOn"`
	PhoneNumber     string    `json:"phoneNumber"`
	OrderStatus     string    `json:"orderStatus"`
	OrderReturned   bool      `json:"orderReturned"`
}
