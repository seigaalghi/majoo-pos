package entities

import (
	"errors"
	"fmt"

	"github.com/seigaalghi/majoo-pos/models"
	"github.com/seigaalghi/majoo-pos/utilities"
	"gorm.io/gorm"
)

func Login(db *gorm.DB, username string, password string) (string, error) {
	var merchant models.Merchant
	db.Where("username = ?", username).Where("password = ?", password).Find(&merchant)
	if merchant.Username == "" {
		return "", errors.New("invalid username or password")
	}
	payload := models.Payload{
		ID:       int(merchant.ID),
		Username: merchant.Username,
	}
	token, err := utilities.GenerateToken(&payload)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	return token, nil
}

func CreateMerchant(db *gorm.DB, merchant *models.MerchantInput) (*models.Merchant, error) {
	var result *models.Merchant
	input := models.Merchant{
		Name:     merchant.Name,
		Password: merchant.Password,
		Username: merchant.Username,
	}
	err := db.Where("username = ?", merchant.Username).First(&result).Error
	if err == nil {
		return nil, errors.New("username already used")
	}
	db.Create(&input)
	return &input, nil
}

func UpdateMerchant(db *gorm.DB, id string, merchant *models.MerchantInput) (*models.Merchant, error) {
	var result *models.Merchant
	input := models.Merchant{
		Name:     merchant.Name,
		Password: merchant.Password,
		Username: merchant.Username,
	}
	err := db.Where("username = ?", merchant.Username).Where("id <> ?", id).First(&result).Error
	if err == nil {
		return nil, errors.New("username already used")
	}
	db.Where("id = ?", id).Find(&result)
	db.Model(&result).Updates(&input)

	return result, nil
}

func DeleteMerchant(db *gorm.DB, id string) error {
	var result *models.Merchant
	err := db.Where("id = ?", id).First(&result).Error
	if err != nil {
		return err
	}
	err = db.Delete(&result).Error
	if err != nil {
		return err
	}
	return nil
}

func GetMerchants(db *gorm.DB) ([]*models.Merchant, error) {
	var result []*models.Merchant
	err := db.Select("name", "username", "id", "CreatedAt").Limit(10).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetMerchant(db *gorm.DB, id string) (*models.Merchant, error) {
	var result *models.Merchant
	err := db.Select("name", "username", "id", "CreatedAt").Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
