package entities

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model

	TotalProduct int  `gorm:"not null" json:"total_product"`
	TotalPrice   int  `gorm:"not null" json:"total_price"`
	UserID       uint `gorm:"not null" json:"user_id"`
	ProductID    uint `gorm:"not null" json:"product_id"`
}
