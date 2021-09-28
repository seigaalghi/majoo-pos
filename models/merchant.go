package models

import "gorm.io/gorm"

type Merchant struct {
	gorm.Model
	Name     string `json:"name" binding:"required" gorm:"type:varchar(50)"`
	Username string `json:"username" binding:"required" gorm:"type:varchar(50)"`
	Password string `json:"-" binding:"required" gorm:"type:varchar(50)"`
}

type MerchantInput struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Payload struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
}
