package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID         uint   `json:"user_id" form:"user_id"`
	Book_DetailsID uint   `json:"book_details_id" form:"book_detail_id"`
	BorrowDate     string `json:"borrow_date" form:"borrow_date"`
	ReturnDate     string `json:"return_date" form:"return_date"`
	Status         string `json:"status" form:"status"`
	Book_Returns   []Book_Return
}
