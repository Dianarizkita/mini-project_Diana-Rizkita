package models

import (
	"gorm.io/gorm"
)

type Book_Details struct {
	gorm.Model
	BookCode string `json:"book_code" form:"book_code"`
	Status   string `json:"status" form:"status"`
	Book     []Book `json:"book" form:"book"`
}
