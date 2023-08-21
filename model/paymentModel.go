package model

import "time"

// Payment is a struct that represents the payment table in the database

type Payment struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	PaymentTypeId int64     `json:"payment_type_id"`
	Logo          string    `json:"logo"`
	CreatedAt     time.Time `json:"created_at"`
	UpdateAt      time.Time `json:"updated_at"`
}

// PaymentType is a struct that represents the payment_type table in the database

type PaymentType struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}
