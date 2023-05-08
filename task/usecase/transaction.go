package usecase

import (
	"errors"
	"fmt"
	"task/lib/database"
	"task/models"
)

func CreateTransaction(transaction *models.Transaction) error {
	// check title cannot be empty
	//if transaction.Status == "" {
	//	return errors.New("Lloan date cannot be empty")
	//}

	//check creator
	//if transaction.DateOfReturn == "" {
	//	return errors.New("date of book cannot be empty")
	//}

	err := database.CreateTransaction(transaction)
	if err != nil {
		return err
	}
	return nil
}

func GetTransaction(id uint) (transaction models.Transaction, err error) {
	transaction, err = database.GetTransaction(id)
	if err != nil {
		fmt.Println("Error getting transaction from database")
		return
	}
	return
}

func GetListTransactions() (transactions []models.Transaction, err error) {
	if err != nil {
		fmt.Println("Error getting transaction from database")
		return
	}
	return
}

func UpdateTransaction(transaction *models.Transaction) (err error) {
	err = database.UpdateTransaction(transaction)
	if err != nil {
		fmt.Println(" Error Updating transaction")
		return
	}

	return
}

func DeleteTransaction(id uint) (err error) {
	transaction := models.Transaction{}
	transaction.ID = id
	err = database.DeleteTransaction(&transaction)
	if err != nil {
		fmt.Println("DeleteBook : error deleting book, err ", err)
		return
	}

	return
}

// ====
func LendBook(Buku_Details, BukuID uint) (book_details *models.Book_Details, err error) {
	// Find the book in our library
	buku_details, err := database.GetBookDetails(Buku_Details)
	if err != nil {
		return nil, err
	}

	// Check if the book is available
	if buku_details.Status != "tidak tersedia" {
		return nil, errors.New("book is not available")
	}

	// Update the book's status
	buku_details.Status = "Telah Dipinjam"
	return
}

func BorrowBook(book_details models.Book_Details, member models.User) error {
	if book_details.Status != "Available" {
		return errors.New("Book is not available for borrowing")
	}

	book_details.Status = "Borrowed"

	return nil
}
