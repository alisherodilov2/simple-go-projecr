package models

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	ProductID   uint     `json:"product_id"`
	Products    Products `gorm:"foreignKey:ProductID"`
}
