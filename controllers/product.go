package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/majoo-pos/entities"
	"github.com/seigaalghi/majoo-pos/models"
	"github.com/seigaalghi/majoo-pos/utilities"
	"gorm.io/gorm"
)

func CreateProduct(ctx *gin.Context) {
	var input *models.ProductInput
	merchantId := ctx.MustGet("id").(int)
	db := ctx.MustGet("db").(*gorm.DB)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utilities.ErrorResponse(ctx, err, 400)
		return
	}
	result, err := entities.CreateProduct(db, merchantId, input)
	if err != nil {
		utilities.ErrorResponse(ctx, err, 400)
		return
	}
	ctx.JSON(201, gin.H{
		"status":  "success",
		"message": "Created successfully",
		"result":  result,
	})
}

func GetProduct(ctx *gin.Context) {
	merchantId := ctx.MustGet("id").(int)
	productId := ctx.Param("id")
	db := ctx.MustGet("db").(*gorm.DB)
	product, err := entities.GetProduct(db, productId, merchantId)
	if err != nil {
		utilities.ErrorResponse(ctx, err, 400)
		return
	}
	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": "Product retrieved successfully",
		"result":  product,
	})
}

func GetProducts(ctx *gin.Context) {
	merchantId := ctx.MustGet("id").(int)
	db := ctx.MustGet("db").(*gorm.DB)
	products, err := entities.GetProducts(db, merchantId)
	if err != nil {
		utilities.ErrorResponse(ctx, err, 400)
		return
	}
	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": "Products retrieved successfully",
		"result":  products,
	})
}

func UpdateProduct(ctx *gin.Context) {
	var input *models.ProductInput
	productId := ctx.Param("id")
	merchantId := ctx.MustGet("id").(int)
	db := ctx.MustGet("db").(*gorm.DB)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utilities.ErrorResponse(ctx, err, 400)
		return
	}
	result, err := entities.UpdateProduct(db, productId, merchantId, input)
	if err != nil {
		utilities.ErrorResponse(ctx, err, 400)
		return
	}
	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": "Updated successfully",
		"result":  result,
	})
}

func DeleteProduct(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	productId := ctx.Param("id")
	merchantId := ctx.MustGet("id").(int)
	err := entities.DeleteProduct(db, productId, merchantId)
	if err != nil {
		utilities.ErrorResponse(ctx, err, 400)
		return

	}
	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": "Product deleted successfully",
		"result":  gin.H{},
	})
}
