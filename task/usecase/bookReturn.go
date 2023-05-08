package usecase

import (
	"fmt"
	"task/lib/database"
	"task/models"
)

func CreateBookReturn(book_return *models.Book_Return) error {
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

//=============

func ReturnBook(TransactionID, Buku_DetailsID uint) (book_return *models.Book_Return, err error) {
	// Find the book in our library
	transaction, err := database.GetTransaction(TransactionID)
	if err != nil {
		return nil, err
	}
	buku_details, err := database.GetBookDetails(Buku_DetailsID)
	if err != nil {
		return nil, err
	}

	// Check if the book is available
	//if transaction.Status != "tidak tersedia" {
	////    return nil, errors.New("transaksi is not available")
	//}

	// Update the book's status
	transaction.Status = "Tersedia"
	buku_details.Status = "Tersedia"
	return
}
