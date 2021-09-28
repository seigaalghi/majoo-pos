package entities

import (
	"github.com/seigaalghi/majoo-pos/models"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, merchantId int, input *models.ProductInput) (*models.Product, error) {
	product := models.Product{
		Name:       input.Name,
		Sku:        input.Sku,
		Image:      input.Image,
		MerchantID: merchantId,
	}
	err := db.Model(&product).Create(&product).Joins("Merchant").Error
	if err != nil {
		return nil, err
	}
	err = db.Joins("Merchant").Find(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func GetProducts(db *gorm.DB, merchantId int) ([]*models.Product, error) {
	var result []*models.Product
	err := db.Where("merchant_id = ?", merchantId).Limit(10).Joins("Merchant").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetProduct(db *gorm.DB, productId string, merchantId int) (*models.Product, error) {
	var result *models.Product
	err := db.Model(&result).Where("products.id = ?", productId).Where("merchant_id = ?", merchantId).Joins("Merchant").First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateProduct(db *gorm.DB, productId string, merchantId int, input *models.ProductInput) (*models.Product, error) {
	var result *models.Product
	product := models.Product{
		Name:  input.Name,
		Image: input.Image,
		Sku:   input.Sku,
	}
	err := db.Where("products.id = ?", productId).Where("merchant_id = ?", merchantId).Joins("Merchant").Find(&result).Error
	if err != nil {
		return nil, err
	}
	err = db.Model(&result).Updates(&product).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteProduct(db *gorm.DB, productId string, merchantId int) error {
	var result *models.Product
	err := db.Where("id = ?", productId).Where("merchant_id = ?", merchantId).First(&result).Error
	if err != nil {
		return err
	}
	err = db.Delete(&result).Error
	if err != nil {
		return err
	}
	return nil
}
