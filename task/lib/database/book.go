package database

import (
	"task/config"
	"task/models"
)

func CreateBook(book *models.Book) error {
	if err := config.DB.Save(book).Error; err != nil {
		return err
	}
	return nil
}

func GetBooks() (interface{}, error) {
	var books []models.Book

	if err := config.DB.Model(&models.Book_Details{}).Preload("Book_Details").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBook(id uint) (book models.Book, err error) {
	book.ID = id
	if err = config.DB.Model(&models.Book_Details{}).Preload("Book_Details").First(&book).Error; err != nil {
		return
	}
	return
}

func UpdateBook(book *models.Book) error {
	if err := config.DB.Updates(book).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBook(book *models.Book) error {
	if err := config.DB.Delete(book).Error; err != nil {
		return err
	}
	return nil
}
