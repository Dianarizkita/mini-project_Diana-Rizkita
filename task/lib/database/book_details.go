package database

import (
	"task/config"
	"task/models"
)

func CreateBookDetails(book_details *models.Book_Details) error {
	if err := config.DB.Save(book_details).Error; err != nil {
		return err
	}
	return nil
}

func GetBooksDetails() (interface{}, error) {
	var books_details []models.Book_Details

	if err := config.DB.Find(&books_details).Error; err != nil {
		return nil, err
	}
	return books_details, nil
}

func GetBookDetails(id uint) (book_details models.Book_Details, err error) {
	book_details.ID = id
	if err = config.DB.First(&book_details).Error; err != nil {
		return
	}
	return
}

func UpdateBookDetails(book_details *models.Book_Details) error {
	if err := config.DB.Updates(book_details).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBookDetails(book_details *models.Book_Details) error {
	if err := config.DB.Delete(book_details).Error; err != nil {
		return err
	}
	return nil
}
