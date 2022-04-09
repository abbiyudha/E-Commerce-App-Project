package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	PhoneNumber  string `json:"phoneNumber" form:"phoneNumber"`
	Address      string `json:"address" form:"address"`
	CartItems    []CartItem
	Transactions []Transaction
}
