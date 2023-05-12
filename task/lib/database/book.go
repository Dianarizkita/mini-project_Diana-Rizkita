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

func GetBooks() (books []models.Book, err error) {
	if err = config.DB.Model(&models.Book{}).Preload("Book_Detailss").Find(&books).Error; err != nil {
		return
	}
	return
}

func GetBook(id uint) (book models.Book, err error) {
	book.ID = id
	if err = config.DB.Model(&models.Book{}).Preload("Book_Detailss").First(&book).Error; err != nil {
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
