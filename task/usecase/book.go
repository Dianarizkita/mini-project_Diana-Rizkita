package usecase

import (
	"errors"
	"fmt"
	"task/lib/database"
	"task/models"
)

func CreateBook(book *models.Book) error {
	// check title cannot be empty
	if book.Title == "" {
		return errors.New("book title cannot be empty")
	}

	//check creator
	if book.Author == "" {
		return errors.New("book Author cannot be empty")
	}

	err := database.CreateBook(book)
	if err != nil {
		return err
	}
	return nil
}

func GetBook(id uint) (book models.Book, err error) {
	book, err = database.GetBook(id)
	if err != nil {
		fmt.Println("Error getting book from database")
		return
	}
	return
}

func GetListBooks() (books []models.Book, err error) {
	if err != nil {
		fmt.Println("GetListBook : Error getting user from database")
		return
	}
	return
}

func UpdateBook(book *models.Book) (err error) {
	err = database.UpdateBook(book)
	if err != nil {
		fmt.Println("UpdateBook : Error Updating Book")
		return
	}

	return
}

func DeleteBook(id uint) (err error) {
	book := models.Book{}
	book.ID = id
	err = database.DeleteBook(&book)
	if err != nil {
		fmt.Println("DeleteBook : error deleting book, err ", err)
		return
	}

	return
}
