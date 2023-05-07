package usecase

import (
	"errors"
	"fmt"
	"task/lib/database"
	"task/models"
)

func CreateBookDetails(book_details *models.Book_Details) error {
	// check code book cannot be empty
	if book_details.BookCode == "" {
		return errors.New("book code cannot be empty")
	}

	//check status
	if book_details.Status == "" {
		return errors.New("status cannot be empty")
	}

	err := database.CreateBookDetails(book_details)
	if err != nil {
		return err
	}
	return nil
}

func GetBookDetails(id uint) (book_details models.Book_Details, err error) {
	book_details, err = database.GetBookDetails(id)
	if err != nil {
		fmt.Println("Error getting book details from database")
		return
	}
	return
}

func GetListBooksDetails() (books_details []models.Book_Details, err error) {
	if err != nil {
		fmt.Println("GetListBook : Error getting user from database")
		return
	}
	return
}

func UpdateBookDetails(book_details *models.Book_Details) (err error) {
	err = database.UpdateBookDetails(book_details)
	if err != nil {
		fmt.Println("UpdateBook : Error Updating Book")
		return
	}

	return
}

func DeleteBookDetails(id uint) (err error) {
	book_details := models.Book_Details{}
	book_details.ID = id
	err = database.DeleteBookDetails(&book_details)
	if err != nil {
		fmt.Println("DeleteBook : error deleting book details, err ", err)
		return
	}

	return
}
