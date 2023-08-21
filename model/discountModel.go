package model

import "time"

// Discount is a struct that represents the discount table in the database

type Discount struct {
	ID                 int64     `json:"id"`
	Qty                int64     `json:"qty"`
	Type               string    `json:"type"`
	Result             int       `json:"result"`
	ExpiredAt          time.Time `json:"expired_at"`
	ExpiredAtFormatted string    `json:"expired_at_formatted"`
	StringFormat       string    `json:"string_format"`
	CreatedAt          time.Time `json:"created_at"`
	UpdateAt           time.Time `json:"updated_at"`
}
