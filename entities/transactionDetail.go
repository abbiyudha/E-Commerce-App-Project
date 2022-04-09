package entities

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model

	TransactionID uint64 `gorm:"not null" json:"transaction_id"`
	ProductID     uint64 `gorm:"not null" json:"product_id"`
	Quantity      int    `gorm:"not null" json:"quantity"`
	Price         int    `gorm:"not null" json:"price"`
}
