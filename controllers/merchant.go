package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/majoo-pos/entities"
	"github.com/seigaalghi/majoo-pos/models"
	"github.com/seigaalghi/majoo-pos/utilities"
	"gorm.io/gorm"
)

func Login(ctx *gin.Context) {
	var input models.Login
	db := ctx.MustGet("db").(*gorm.DB)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utilities.ErrorResponse(ctx, err, 400)
		return
	}
	token, err := entities.Login(db, input.Username, input.Password)
	if err != nil {
		utilities.ErrorResponse(ctx, err, 401)
		return
	}
	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": "Logged in successfully",
		"result": gin.H{
			"token": token,
		},
	})
}

func CreateMerchant(ctx *gin.Context) {
	var input *models.MerchantInput
	db := ctx.MustGet("db").(*gorm.DB)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utilities.ErrorResponse(ctx, err, 400)
		return
	}

	merchant, err := entities.CreateMerchant(db, input)

	if err != nil {
		utilities.ErrorResponse(ctx, err, 400)
		return
	}

	ctx.JSON(201, gin.H{
		"status":  "success",
		"message": "Merchant created successfully",
		"result":  merchant,
	})
}

func UpdateMerchant(ctx *gin.Context) {
	var input *models.MerchantInput
	db := ctx.MustGet("db").(*gorm.DB)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utilities.ErrorResponse(ctx, err, 400)
		return
	}
	id := ctx.Param("id")
	merchant, err := entities.UpdateMerchant(db, id, input)
	if err != nil {
		utilities.ErrorResponse(ctx, err, 400)
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": "Merchant updated successfully",
		"result":  merchant,
	})
}

func DeleteMerchant(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	id := ctx.Param("id")
	err := entities.DeleteMerchant(db, id)
	if err != nil {
		utilities.ErrorResponse(ctx, err, 400)
		return

	}
	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": "Merchant deleted successfully",
		"result":  gin.H{},
	})
}

func GetMerchants(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	merchants, err := entities.GetMerchants(db)
	if err != nil {
		utilities.ErrorResponse(ctx, err, 404)
		return
	}
	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": "Merchants retrieved successfully",
		"result":  merchants,
	})
}

func GetMerchant(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	id := ctx.Param("id")
	merchants, err := entities.GetMerchant(db, id)
	if err != nil {
		utilities.ErrorResponse(ctx, err, 404)
		return
	}
	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": "Merchant retrieved successfully",
		"result":  merchants,
	})
}
