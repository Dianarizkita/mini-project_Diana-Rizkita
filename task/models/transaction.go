package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID         uint      `json:"user_id" form:"user_id"`
	Book_DetailsID uint      `json:"book_id" form:"book_id"`
	LoanDate       time.Time `json:"loan_date"`
	DateOfReturn   time.Time `json:"date_of_return"`
	Status         string    `json:"status" form:"status"`
	Book_Returns   []Book_Return
}
