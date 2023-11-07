package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Email      string `json:"email"`
	Status     bool   `json:"status"`
	IsArchived bool   `json:"is_archived"`
}

type OrderItem struct {
	gorm.Model
	OrderID      uint `json:"order_id"`
	ItemID       uint `json:"item_id"`
	ItemQuantity uint `json:"item_quantity"`

	Title         string `json:"title"`
	CategoryId    uint   `json:"category_id"`
	SubCategoryId uint   `json:"sub_category_id"`
	Img           string `json:"img"`
	Type          string `json:"type"`
}
