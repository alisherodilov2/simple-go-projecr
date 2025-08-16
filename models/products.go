package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Image       string     `json:"image"`
	Comments    []Comments `gorm:"foreignKey:ProductID"`
}
