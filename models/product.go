package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string   `json:"name" binding:"required" gorm:"type:varchar(70)"`
	Sku        int      `json:"sku" binding:"required" gorm:"type:int(50)"`
	Image      string   `json:"image" gorm:"type:varchar(255)"`
	MerchantID int      `json:"merchant_id" binding:"required"`
	Merchant   Merchant `json:"merchant" gorm:"foreignKey:MerchantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ProductInput struct {
	Name       string `json:"name" binding:"required"`
	Sku        int    `json:"sku" binding:"required" `
	Image      string `json:"image"`
	MerchantID int    `json:"merchant_id"`
}
