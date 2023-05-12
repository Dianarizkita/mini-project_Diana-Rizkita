package database

import (
	"task/config"
	"task/models"
)

func CreateBookReturn(book_return *models.Book_Return) error {
	if err := config.DB.Save(book_return).Error; err != nil {
		return err
	}
	return nil
}

func GetBooksReturn() (books_return []models.Book_Return, err error) {
	if err = config.DB.Find(&books_return).Error; err != nil {
		return
	}
	return
}

func GetBookReturn(id uint) (book_return models.Book_Return, err error) {
	book_return.ID = id
	if err = config.DB.First(&book_return).Error; err != nil {
		return
	}
	return
}

func UpdateBookReturn(book_return *models.Book_Return) error {
	if err := config.DB.Updates(book_return).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBookReturn(book_return *models.Book_Return) error {
	if err := config.DB.Delete(book_return).Error; err != nil {
		return err
	}
	return nil
}
