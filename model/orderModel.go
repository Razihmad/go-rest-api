package model

import "time"

// Order is a struct that represents the order table in the database

type Order struct {
	ID             int64     `json:"id"`
	CashierID      int64     `json:"cashier_id"`
	PaymentTypesId int64     `json:"payment_types_id"`
	TotalPrice     int       `json:"total_price"`
	TotalPaid      int       `json:"total_paid"`
	TotalReturn    int       `json:"total_return"`
	RecieptId      int64     `json:"reciept_id"`
	IsDownloaded   bool      `json:"is_downloaded"`
	ProductId      int64     `json:"product_id"`
	Quantity       int       `json:"quantity"`
	CreatedAt      time.Time `json:"created_at"`
	UpdateAt       time.Time `json:"updated_at"`
}
