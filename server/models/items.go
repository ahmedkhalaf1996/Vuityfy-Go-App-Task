package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Title string `json:"title"`
}

type SubCategory struct {
	gorm.Model
	Title      string `json:"title"`
	CategoryId uint   `json:"category_id"`
}

type Item struct {
	gorm.Model
	Title         string `json:"title"`
	CategoryId    uint   `json:"category_id"`
	SubCategoryId uint   `json:"sub_category_id"`
	Img           string `json:"img"`
	Type          string `json:"type"`
}
