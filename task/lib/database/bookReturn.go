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

func GetBookReturns() (interface{}, error) {
	var book_returns []models.Book_Return

	if err := config.DB.Find(&book_returns).Error; err != nil {
		return nil, err
	}
	return book_returns, nil
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
