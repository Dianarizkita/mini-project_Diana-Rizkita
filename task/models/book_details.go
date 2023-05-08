package models

import (
	"gorm.io/gorm"
)

type Book_Details struct {
	gorm.Model
	BookID       uint   `json:"book_id" form:"book_id"`
	BookCode     string `json:"book_code" form:"book_code"`
	Status       string `json:"status" form:"status"`
	Transactions []Transaction
}
