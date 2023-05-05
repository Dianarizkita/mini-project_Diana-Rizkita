package usecase

import (
	"errors"
	"fmt"
	"task/lib/database"
	"task/models"
)

func CreateBookReturn(book_return *models.Book_Return) error {
	// check title cannot be empty
	if book_return.DateOfReturn == "" {
		return errors.New("date of return cannot be empty")
	}

	//check creator
	if book_return.Penalty == "" {
		return errors.New("penalty cannot be empty")
	}

	err := database.CreateBookReturn(book_return)
	if err != nil {
		return err
	}
	return nil
}

func GetBookReturn(id uint) (book_return models.Book_Return, err error) {
	book_return, err = database.GetBookReturn(id)
	if err != nil {
		fmt.Println("Error getting book return from database")
		return
	}
	return
}

func GetListBookReturns() (book_returns []models.Book_Return, err error) {
	if err != nil {
		fmt.Println("GetListBookReturn : Error getting book return from database")
		return
	}
	return
}

func UpdateBookReturn(book_return *models.Book_Return) (err error) {
	err = database.UpdateBookReturn(book_return)
	if err != nil {
		fmt.Println("UpdateBook : Error Updating Book")
		return
	}

	return
}

func DeleteBookReturn(id uint) (err error) {
	book_return := models.Book_Return{}
	book_return.ID = id
	err = database.DeleteBookReturn(&book_return)
	if err != nil {
		fmt.Println("DeleteBook : error deleting book, err ", err)
		return
	}

	return
}
