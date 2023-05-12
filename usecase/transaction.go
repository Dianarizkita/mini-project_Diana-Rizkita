package usecase

import (
	"fmt"
	"task/lib/database"
	"task/models"
)

func CreateTransaction(transaction *models.Transaction) error {
	// check title cannot be empty
	//if book.Title == "" {
	//		return errors.New("book title cannot be empty")
	//	}

	//check creator
	//	if book.Author == "" {
	//		return errors.New("book Author cannot be empty")
	//	}

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
	transactions, err = database.GetTransactions()
	if err != nil {
		fmt.Println("GetListTransaction : Error getting transaction from database")
		return
	}
	return
}

func UpdateTransaction(transaction *models.Transaction) (err error) {
	err = database.UpdateTransaction(transaction)
	if err != nil {
		fmt.Println("UpdateTransaction : Error Updating transaction")
		return
	}

	return
}

func DeleteTransaction(id uint) (err error) {
	transaction := models.Transaction{}
	transaction.ID = id
	err = database.DeleteTransaction(&transaction)
	if err != nil {
		fmt.Println("DeleteTransaction : error deleting transaction, err ", err)
		return
	}

	return
}
