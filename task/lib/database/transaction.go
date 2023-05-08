package database

import (
	"task/config"
	"task/models"
)

func CreateTransaction(transaction *models.Transaction) error {
	if err := config.DB.Save(transaction).Error; err != nil {
		return err
	}
	return nil
}

func GetTransactions() (interface{}, error) {
	var transactions []models.Transaction

	if err := config.DB.Model(&models.Transaction{}).Preload("Book_Returns").Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func GetTransaction(id uint) (transaction models.Transaction, err error) {
	transaction.ID = id
	if err = config.DB.Model(&models.Transaction{}).Preload("Book_Returns").First(&transaction).Error; err != nil {
		return
	}
	return
}

func UpdateTransaction(transaction *models.Transaction) error {
	if err := config.DB.Updates(transaction).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTransaction(transaction *models.Transaction) error {
	if err := config.DB.Delete(transaction).Error; err != nil {
		return err
	}
	return nil
}
