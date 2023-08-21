package model

import "time"

// Product is a struct that represents the product table in the database

type Product struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	Sku              string    `json:"sku"`
	Price            int64     `json:"price"`
	Stock            int64     `json:"stock"`
	CategoryId       int64     `json:"category_id"`
	DiscountId       int64     `json:"discount_id"`
	Image            string    `json:"image"`
	TotalFinalPrice  int64     `json:"total_final_price"`
	TotalNormalPrice int64     `json:"total_normal_price"`
	CreatedAt        time.Time `json:"created_at"`
	UpdateAt         time.Time `json:"updated_at"`
}
