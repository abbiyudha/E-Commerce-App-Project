package entities

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	TotalPrice int  `gorm:"type:int;not null" json:"total_price"`
	UserID     uint `gorm:"not null" json:"user_id"`

	TransactionDetail []TransactionDetail
}
